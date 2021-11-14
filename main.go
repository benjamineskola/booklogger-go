package main

import (
	"booklogger/controllers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=ben dbname=ben port=5432 sslmode=disable"
	}

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	server := gin.Default()
	router := server.Group("/")

	if os.Getenv("AUTH_USER") != "" && os.Getenv("AUTH_PASSWORD") != "" {
		router = server.Group("/", gin.BasicAuth(gin.Accounts{
			os.Getenv("AUTH_USER"): os.Getenv("AUTH_PASSWORD"),
		}))
	}

	router.GET("/books.json", func(c *gin.Context) { controllers.BookList(c, database) })
	router.GET("/books/:slug.json", func(c *gin.Context) { controllers.BookBySlug(c, database) })
	router.GET("/authors.json", func(c *gin.Context) { controllers.AuthorList(c, database) })
	log.Fatal(server.Run())
}
