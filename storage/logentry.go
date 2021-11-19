package storage

import (
	"booklogger/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetAllLogEntries(db *gorm.DB) (logentries *[]models.LogEntry, err error) {
	err = db.Model(&models.LogEntry{}).
		Preload(clause.Associations).
		Preload("Book.FirstAuthor").
		Find(&logentries).Error

	return
}
