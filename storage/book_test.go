package storage_test

import (
	"booklogger/models"
	"booklogger/storage"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestBook(t *testing.T) { //nolint:funlen,paralleltest
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

	err = database.AutoMigrate(&models.Book{}, &models.Author{})
	if err != nil {
		panic(err)
	}

	t.Run("get data when empty", func(t *testing.T) { //nolint:paralleltest
		books, err := storage.GetAllBooks(database)
		assert.Nil(err)
		assert.Len(*books, 0)
	})

	t.Run("get data when nonempty", func(t *testing.T) { //nolint:paralleltest
		author := *models.NewAuthor("Agatha Christie")
		err := database.Create(&author).Error
		assert.Nil(err)

		book1 := *models.NewBook("The Mysterious Affair at Styles")
		book1.FirstAuthor = author
		book1.Slug = "christie-mysterious-affair-styles"
		err = database.Create(&book1).Error
		assert.Nil(err)

		book2 := *models.NewBook("Murder at the Vicarage")
		book2.FirstAuthor = author
		err = database.Create(&book2).Error
		assert.Nil(err)

		books, err := storage.GetAllBooks(database)
		assert.Nil(err)
		assert.Len(*books, 2)
	})

	t.Run("get data for an individual book", func(t *testing.T) { //nolint:paralleltest
		_, err := storage.GetBookBySlug(database, "christie-mysterious-affair-styles")
		assert.Nil(err)
	})

	t.Run("error if book does not exist", func(t *testing.T) { //nolint:paralleltest
		result, err := storage.GetBookBySlug(database, "no-such-book")
		assert.NotNil(err, "should have returned an error but returned "+result.String())
	})

	t.Cleanup(func() {
		database.Exec(
			"drop table library_author,library_book,library_bookauthor,library_book_editions cascade;",
		)
	})
}
