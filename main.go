package main

import (
	"booklogger/controllers"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func fixParamMiddleware(ctx *gin.Context) {
	for n, param := range ctx.Params {
		if strings.HasSuffix(param.Key, ".json") {
			ctx.Params[n] = gin.Param{
				Key:   strings.TrimSuffix(param.Key, ".json"),
				Value: strings.TrimSuffix(param.Value, ".json"),
			}
			ctx.Params = append(ctx.Params, gin.Param{Key: "format", Value: "json"})
		}
	}

	ctx.Next()
}

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
	server.Use(fixParamMiddleware)
	router := server.Group("/")

	if os.Getenv("AUTH_USER") != "" && os.Getenv("AUTH_PASSWORD") != "" {
		router = server.Group("/", gin.BasicAuth(gin.Accounts{
			os.Getenv("AUTH_USER"): os.Getenv("AUTH_PASSWORD"),
		}))
	}

	handleWithDB := func(fun func(*gin.Context, *gorm.DB)) func(*gin.Context) {
		return func(ctx *gin.Context) {
			fun(ctx, database)
		}
	}

	router.GET("/books.json", handleWithDB(controllers.BookList))
	router.GET("/books/:slug.json", handleWithDB(controllers.BookBySlug))
	router.GET("/authors.json", handleWithDB(controllers.AuthorList))
	log.Fatal(server.Run())
}
