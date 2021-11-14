package controllers

import (
	"booklogger/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthorList(c *gin.Context, db *gorm.DB) {
	authors := storage.GetAllAuthors(db)

	c.JSON(http.StatusOK, authors)
}
