package controllers

import (
	"booklogger/storage"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

func AuthorList(resp http.ResponseWriter, req *http.Request, db *gorm.DB) {
	authors := storage.GetAllAuthors(db)

	if err := json.NewEncoder(resp).Encode(authors); err != nil {
		panic(err)
	}
}
