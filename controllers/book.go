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
