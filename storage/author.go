package storage

import (
	"booklogger/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetAllAuthors(db *gorm.DB) (authors *[]models.Author, err error) {
	result := db.Model(&models.Author{}).
		Preload(clause.Associations).
		Preload("FirstAuthoredBooks.FirstAuthor").
		// Preload("FirstAuthoredBooks.AdditionalAuthors").
		Preload("AdditionalAuthoredBooks.FirstAuthor").
		// Preload("AdditionalAuthoredBooks.AdditionalAuthors").
		Find(&authors)
	if result.Error != nil { // notest
		err = result.Error
	}

	return
}

func GetAuthorBySlug(db *gorm.DB, slug string) (*models.Author, error) {
	var author models.Author

	result := db.Model(&models.Author{}).
		Preload(clause.Associations).
		Where("slug = ?", slug).
		First(&author)
	if result.Error != nil {
		return nil, result.Error
	}

	return &author, result.Error
}
