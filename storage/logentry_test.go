package storage_test

import (
	"booklogger/models"
	"booklogger/storage"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestLogEntry(t *testing.T) { //nolint:funlen,paralleltest
	assert := assert.New(t)

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
		entries, err := storage.GetAllLogEntries(database)
		assert.Nil(err)
		assert.Len(*entries, 0)
	})

	t.Run("get data when nonempty", func(t *testing.T) { //nolint:paralleltest
		author := *models.NewAuthor("Agatha Christie")
		err := database.Create(&author).Error
		assert.Nil(err)

		book := models.NewBook("The Mysterious Affair at Styles")
		book.FirstAuthor = author
		book.Slug = "christie-mysterious-affair-styles"
		err = database.Create(book).Error
		assert.Nil(err)

		book, _ = storage.GetBookBySlug(database, "christie-mysterious-affair-styles")

		entry := models.LogEntry{Book: *book}
		err = database.Create(&entry).Error
		assert.Nil(err)

		entries, _ := storage.GetAllLogEntries(database)
		assert.Len(*entries, 1)
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
