package controllers

import (
	"booklogger/handlers"
	"booklogger/storage"
	"encoding/json"
	"net/http"
)

func AuthorList(ctx *handlers.Context) (status int, result []byte) {
	authors := storage.GetAllAuthors(ctx.App.DB)

	result, err := json.Marshal(authors)
	if err != nil {
		result = []byte(err.Error())
		status = http.StatusInternalServerError
	}

	return
}
