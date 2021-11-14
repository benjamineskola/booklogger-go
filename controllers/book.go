package controllers

import (
	"booklogger/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BookList(c *gin.Context, db *gorm.DB) {
	books := storage.GetAllBooks(db)

	c.JSON(http.StatusOK, books)
}

func BookBySlug(ctx *gin.Context, database *gorm.DB) {
	if slug := ctx.Param("slug"); slug != "" {
		if book, err := storage.GetBookBySlug(database, slug); err == nil {
			ctx.JSON(http.StatusOK, book)
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err})
		}
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "no slug given"})
	}
}
