package controllers

import (
	"book-store/models"

	"github.com/gin-gonic/gin"
)

func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(200, gin.H{"data": books})
}
