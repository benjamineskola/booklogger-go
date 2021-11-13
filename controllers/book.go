package controllers

import (
	"booklogger/storage"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

func BookList(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	books := storage.GetAllBooks(db)

	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		panic(err)
	}
}
