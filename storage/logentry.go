package storage

import (
	"booklogger/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetAllLogEntries(db *gorm.DB) (logentries *[]models.LogEntry, err error) {
	result := db.Model(&models.LogEntry{}).
		Preload(clause.Associations).
		Preload("Book.FirstAuthor").
		Find(&logentries)
	if result.Error != nil { // notest
		err = result.Error
	}

	return
}
