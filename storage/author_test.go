package storage_test

import (
	"booklogger/models"
	"booklogger/storage"
	"os"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestAuthor(t *testing.T) { //nolint:paralleltest
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

	err = database.AutoMigrate(&models.Book{}, &models.Author{})
	if err != nil {
		panic(err)
	}

	t.Run("get data when empty", func(t *testing.T) { //nolint:paralleltest
		books := *storage.GetAllAuthors(database)
		if len(books) != 0 {
			t.Fatal("Authors list should be empty, not", len(books), books)
		}
	})

	t.Run("get data when nonempty", func(t *testing.T) { //nolint:paralleltest
		author := *models.NewAuthor("Agatha Christie")
		result := database.Create(&author)
		if result.Error != nil {
			t.Fatal(err)
		}

		authors := *storage.GetAllAuthors(database)
		if len(authors) != 1 {
			t.Fatal("Books list should be 1, not", len(authors))
		}
	})

	t.Cleanup(func() {
		database.Exec(
			"drop table library_author,library_book,library_bookauthor,library_book_editions cascade;",
		)
	})
}
