package controllers

import (
	h "booklogger/http"
	"booklogger/storage"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func BookList(resp http.ResponseWriter, req *http.Request, db *gorm.DB) {
	books := storage.GetAllBooks(db)
	h.JSONResponse(resp, books)
}

func BookBySlug(resp http.ResponseWriter, req *http.Request, database *gorm.DB) {
	if slug := mux.Vars(req)["slug"]; slug != "" {
		book, err := storage.GetBookBySlug(database, slug)
		if err == nil {
			h.JSONResponse(resp, book)
		} else {
			h.JSONError(resp, http.StatusNotFound, err.Error())
		}
	} else {
		h.JSONError(resp, http.StatusBadRequest, "no slug given")
	}
}
