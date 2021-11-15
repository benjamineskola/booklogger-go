package controllers

import (
	"booklogger/storage"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func BookList(resp http.ResponseWriter, req *http.Request, db *gorm.DB) {
	books := storage.GetAllBooks(db)
	if err := json.NewEncoder(resp).Encode(books); err != nil {
		panic(err)
	}
}

func BookBySlug(resp http.ResponseWriter, req *http.Request, database *gorm.DB) {
	if slug := mux.Vars(req)["slug"]; slug != "" { //nolint:nestif
		book, err := storage.GetBookBySlug(database, slug)
		if err == nil {
			if jsonErr := json.NewEncoder(resp).Encode(book); jsonErr != nil {
				panic(jsonErr)
			}
		} else {
			resp.WriteHeader(http.StatusNotFound)
			if jsonErr := json.NewEncoder(resp).Encode(map[string]string{"error": err.Error()}); jsonErr != nil {
				panic(jsonErr)
			}
		}
	} else {
		resp.WriteHeader(http.StatusBadRequest)
		if jsonErr := json.NewEncoder(resp).Encode(map[string]string{"error": "no slug given"}); jsonErr != nil {
			panic(jsonErr)
		}
	}
}
