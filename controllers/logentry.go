package controllers

import (
	"booklogger/handlers"
	"booklogger/storage"
	"encoding/json"
	"net/http"
)

func LogEntryList(ctx *handlers.Context) (status int, result []byte) {
	entries, err := storage.GetAllLogEntries(ctx.App.DB)

	if err == nil {
		result, err = json.Marshal(entries)
		if err != nil {
			result = []byte(err.Error())
			status = http.StatusInternalServerError
		}
	} else {
		result = []byte(err.Error())
	}

	return
}
