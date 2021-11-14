package storage

import (
	"booklogger/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetAllBooks(db *gorm.DB) *[]models.Book {
	var books []models.Book

	result := db.Model(&models.Book{}).Preload(clause.Associations).Find(&books)
	if result.Error != nil {
		panic(result.Error)
	}

	return &books
}

func GetBookBySlug(db *gorm.DB, slug string) (*models.Book, error) {
	var book models.Book

	result := db.Model(&models.Book{}).
		Preload(clause.Associations).
		Where("slug = ?", slug).
		First(&book)
	if result.Error != nil {
		return nil, result.Error
	}

	return &book, result.Error
}
