package controllers

import (
	"booklogger/handlers"
	"booklogger/storage"
	"encoding/json"
	"net/http"
)

func BookList(ctx *handlers.Context) (status int, result []byte) {
	books, err := storage.GetAllBooks(ctx.App.DB)

	if err == nil {
		result, err = json.Marshal(books)
		if err != nil {
			result = []byte(err.Error())
			status = http.StatusInternalServerError
		}
	} else {
		var err error
		result, err = json.Marshal([]byte(err.Error()))
		if err != nil {
			panic(err)
		}
		status = http.StatusInternalServerError
	}

	return
}

func BookBySlug(ctx *handlers.Context) (status int, result []byte) {
	slug := ctx.Vars["slug"]
	book, err := storage.GetBookBySlug(ctx.App.DB, slug)

	if err == nil {
		var jsonErr error
		result, jsonErr = json.Marshal(book)

		if jsonErr != nil {
			result = []byte(jsonErr.Error())
			status = http.StatusInternalServerError
		}
	} else {
		result = []byte(err.Error())
		status = http.StatusNotFound
	}

	return
}
