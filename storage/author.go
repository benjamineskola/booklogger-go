package storage

import (
	"booklogger/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetAllAuthors(db *gorm.DB) *[]models.Author {
	var authors []models.Author

	result := db.Model(&models.Author{}).
		Preload(clause.Associations).
		Preload("FirstAuthoredBooks.FirstAuthor").
		// Preload("FirstAuthoredBooks.AdditionalAuthors").
		Preload("AdditionalAuthoredBooks.FirstAuthor").
		// Preload("AdditionalAuthoredBooks.AdditionalAuthors").
		Find(&authors)
	if result.Error != nil {
		panic(result.Error)
	}

	return &authors
}
