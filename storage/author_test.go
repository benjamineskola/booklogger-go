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

func TestAuthor(t *testing.T) { //nolint:paralleltest
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
		authors, err := storage.GetAllAuthors(database)
		assert.Nil(err)
		assert.Len(*authors, 0)
	})

	t.Run("get data when nonempty", func(t *testing.T) { //nolint:paralleltest
		author := *models.NewAuthor("Agatha Christie")
		author.Slug = "christie-a"

		err := database.Create(&author).Error
		assert.Nil(err)

		authors, err := storage.GetAllAuthors(database)
		assert.Nil(err)
		assert.Len(*authors, 1)
	})

	t.Run("get data about individual author", func(t *testing.T) { //nolint:paralleltest
		author, err := storage.GetAuthorBySlug(database, "christie-a")
		assert.Nil(err)
		assert.Equal("Agatha Christie", author.DisplayName())
	})

	t.Cleanup(func() {
		database.Exec(
			"drop table library_author,library_book,library_bookauthor,library_book_editions cascade;",
		)
	})
}
