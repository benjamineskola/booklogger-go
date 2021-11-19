package storage_test

import (
	"booklogger/models"
	"booklogger/storage"
	"os"
	"testing"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestLogEntry(t *testing.T) { //nolint:cyclop,funlen,paralleltest
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = os.ExpandEnv(
			"host=localhost user=$USER dbname=${USER}_test port=5432 sslmode=disable",
		)
	}

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = database.AutoMigrate(&models.Book{}, &models.Author{}, &models.LogEntry{})
	if err != nil {
		panic(err)
	}

	t.Run("get data when empty", func(t *testing.T) { //nolint:paralleltest
		entries, _ := storage.GetAllLogEntries(database)
		if len(*entries) != 0 {
			t.Fatal("Logentries list should be empty, not", len(*entries), entries)
		}
	})

	t.Run("get data when nonempty", func(t *testing.T) { //nolint:paralleltest
		author := *models.NewAuthor("Agatha Christie")
		result := database.Create(&author)
		if result.Error != nil {
			t.Fatal(result.Error)
		}

		book := models.NewBook("The Mysterious Affair at Styles")
		book.FirstAuthor = author
		book.Slug = "christie-mysterious-affair-styles"
		result = database.Create(book)
		if result.Error != nil {
			t.Fatal(result.Error)
		}

		book, _ = storage.GetBookBySlug(database, "christie-mysterious-affair-styles")

		entry := models.LogEntry{Book: *book}
		result = database.Create(&entry)
		if result.Error != nil {
			t.Fatal(result.Error)
		}

		entries, _ := storage.GetAllLogEntries(database)
		if len(*entries) != 1 {
			t.Fatal("Logentries list should have 1 item, not", len(*entries), entries)
		}
	})

	t.Run("get item by year", func(t *testing.T) { //nolint:paralleltest
		allEntries, _ := storage.GetAllLogEntries(database)
		entry := (*allEntries)[0]
		start, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
		entry.StartAt(&start)
		entry.EndNow()
		database.Save(&entry)

		year := time.Now().Year()

		currentEntries, err := storage.GetLogEntriesByYear(database, year)
		if err != nil {
			t.Fatal(err)
		}

		if len(*currentEntries) != 1 {
			t.Fatal("Logentries list should have 1 item, not", len(*currentEntries), currentEntries)
		}

		prevEntries, err := storage.GetLogEntriesByYear(database, year-1)
		if err != nil {
			t.Fatal(err)
		}

		if len(*prevEntries) != 0 {
			t.Fatal("Logentries list should have no items, not", len(*prevEntries), prevEntries)
		}
	})

	t.Cleanup(func() {
		database.Exec(
			"drop table library_author,library_book,library_bookauthor,library_book_editions,library_logentry cascade;",
		)
	})
}
