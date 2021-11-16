package controllers

import (
	h "booklogger/http"
	"booklogger/storage"
	"net/http"

	"gorm.io/gorm"
)

func AuthorList(resp http.ResponseWriter, req *http.Request, db *gorm.DB) {
	authors := storage.GetAllAuthors(db)
	h.JSONResponse(resp, authors)
}
