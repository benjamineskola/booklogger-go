package storage

import (
	"booklogger/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetAllAuthors(db *gorm.DB) (authors *[]models.Author, err error) {
	err = db.Model(&models.Author{}).
		Preload(clause.Associations).
		Preload("FirstAuthoredBooks.FirstAuthor").
		// Preload("FirstAuthoredBooks.AdditionalAuthors").
		Preload("AdditionalAuthoredBooks.FirstAuthor").
		// Preload("AdditionalAuthoredBooks.AdditionalAuthors").
		Find(&authors).Error

	return
}

func GetAuthorBySlug(db *gorm.DB, slug string) (author *models.Author, err error) {
	err = db.Model(&models.Author{}).
		Preload(clause.Associations).
		Where("slug = ?", slug).
		First(&author).Error

	return
}
