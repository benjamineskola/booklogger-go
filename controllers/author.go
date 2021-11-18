package controllers

import (
	"booklogger/handlers"
	"booklogger/storage"
	"encoding/json"
	"net/http"
)

func AuthorList(ctx *handlers.Context) (status int, result []byte) {
	authors, err := storage.GetAllAuthors(ctx.App.DB)

	if err == nil {
		result, err = json.Marshal(authors)
		if err != nil {
			result = []byte(err.Error())
			status = http.StatusInternalServerError
		}
	} else {
		result = []byte(err.Error())
	}

	return
}

func AuthorBySlug(ctx *handlers.Context) (status int, result []byte) {
	slug := ctx.Vars["slug"]
	author, err := storage.GetAuthorBySlug(ctx.App.DB, slug)

	if err == nil {
		var jsonErr error
		result, jsonErr = json.Marshal(author)

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
