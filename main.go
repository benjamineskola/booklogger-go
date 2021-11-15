package main

import (
	"booklogger/controllers"
	"crypto/subtle"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func basicAuthMiddleware(next http.Handler) http.Handler {
	expectedUser := os.Getenv("AUTH_USER")
	expectedPass := os.Getenv("AUTH_PASSWORD")

	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if expectedUser == "" && expectedPass == "" { //nolint:nestif
			next.ServeHTTP(resp, req)
		} else {
			sentUser, sentPass, ok := req.BasicAuth()
			if !ok || subtle.ConstantTimeCompare([]byte(sentUser), []byte(expectedUser)) != 1 ||
				subtle.ConstantTimeCompare([]byte(sentPass), []byte(expectedPass)) != 1 {
				resp.Header().Set("WWW-Authenticate", `Basic realm="booklogger"`)
				resp.WriteHeader(http.StatusForbidden)
				_, err := resp.Write([]byte("Unauthorised.\n"))
				if err != nil {
					panic(err)
				}
			} else {
				next.ServeHTTP(resp, req)
			}
		}
	})
}

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=ben dbname=ben port=5432 sslmode=disable"
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	router.Use(basicAuthMiddleware)

	router.HandleFunc(
		"/books.json",
		func(w http.ResponseWriter, r *http.Request) { controllers.BookList(w, r, database) },
	)
	router.HandleFunc(
		"/books/{slug}.json",
		func(w http.ResponseWriter, r *http.Request) { controllers.BookBySlug(w, r, database) },
	)
	router.HandleFunc("/authors.json",
		func(w http.ResponseWriter, r *http.Request) { controllers.AuthorList(w, r, database) },
	)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
