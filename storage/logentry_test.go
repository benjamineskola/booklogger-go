package storage_test

import (
	"booklogger/models"
	"booklogger/storage"
	"os"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestLogEntry(t *testing.T) { //nolint:paralleltest
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
			t.Fatal(err)
		}

		book := models.NewBook("The Mysterious Affair at Styles")
		book.FirstAuthor = author
		book.Slug = "christie-mysterious-affair-styles"
		result = database.Create(book)
		if result.Error != nil {
			t.Fatal(err)
		}

		book, _ = storage.GetBookBySlug(database, "christie-mysterious-affair-styles")

		entry := models.LogEntry{Book: *book}
		result = database.Create(&entry)
		if result.Error != nil {
			t.Fatal(err)
		}

		entries, _ := storage.GetAllLogEntries(database)
		if len(*entries) != 1 {
			t.Fatal("Logentries list should be empty, not", len(*entries), entries)
		}
	})

	t.Cleanup(func() {
		database.Exec(
			"drop table library_author,library_book,library_bookauthor,library_book_editions,library_logentry cascade;",
		)
	})
}
