package storage_test

import (
	"booklogger/models"
	"booklogger/storage"
	"os"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestBook(t *testing.T) { //nolint:paralleltest
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=ben dbname=ben_test port=5432 sslmode=disable"
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
		books := *storage.GetAllBooks(database)
		if len(books) != 0 {
			t.Fatal("Books list should be empty, not", len(books), books)
		}
	})

	t.Run("get data when nonempty", func(t *testing.T) { //nolint:paralleltest
		author := *models.NewAuthor("Agatha Christie")
		result := database.Create(&author)
		if result.Error != nil {
			t.Fatal(err)
		}

		book1 := *models.NewBook("The Mysterious Affair at Styles")
		book1.FirstAuthor = author
		result = database.Create(&book1)
		if result.Error != nil {
			t.Fatal(err)
		}

		book2 := *models.NewBook("Murder at the Vicarage")
		book2.FirstAuthor = author
		result = database.Create(&book2)
		if result.Error != nil {
			t.Fatal(err)
		}

		books := *storage.GetAllBooks(database)
		if len(books) != 2 {
			t.Fatal("Books list should be 2, not", len(books))
		}
	})

	t.Cleanup(func() {
		database.Exec(
			"drop table library_author,library_book,library_bookauthor,library_book_editions cascade;",
		)
	})
}
