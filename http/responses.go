package http

import (
	"encoding/json"
	"net/http"
)

func JSONResponse(resp http.ResponseWriter, result interface{}) {
	if err := json.NewEncoder(resp).Encode(result); err != nil {
		panic(err)
	}
}

func JSONError(resp http.ResponseWriter, status int, message string) {
	resp.WriteHeader(status)

	if err := json.NewEncoder(resp).Encode(map[string]string{"error": message}); err != nil {
		panic(err)
	}
}
