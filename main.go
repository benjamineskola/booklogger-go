package main

import (
	"booklogger/controllers"
	"log"
	"net/http"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	port := ":8000"
	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		dsn = "host=localhost user=ben password= dbname=ben port=5432 sslmode=disable"
		port = ":80"
	}

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	http.HandleFunc(
		"/books.json",
		func(w http.ResponseWriter, r *http.Request) { controllers.BookList(w, r, database) },
	)
	log.Fatal(http.ListenAndServe(port, nil))
}
