package main

import (
	"booklogger/controllers"
	"booklogger/handlers"
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

	if expectedUser == "" && expectedPass == "" {
		return http.HandlerFunc(next.ServeHTTP)
	}

	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
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

	app := handlers.InitApp(router, database)
	app.AddJSONRoute("/books.json", controllers.BookList)
	app.AddJSONRoute("/books/{slug}.json", controllers.BookBySlug)
	app.AddJSONRoute("/authors.json", controllers.AuthorList)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
