package storage

import (
	"booklogger/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetAllBooks(db *gorm.DB) (books *[]models.Book, err error) {
	err = db.Model(&models.Book{}).Preload(clause.Associations).Find(&books).Error

	return
}

func GetBookBySlug(db *gorm.DB, slug string) (book *models.Book, err error) {
	err = db.Model(&models.Book{}).
		Preload(clause.Associations).
		Where("slug = ?", slug).
		First(&book).Error

	return
}
