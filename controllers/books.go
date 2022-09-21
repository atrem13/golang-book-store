package controllers


import (
	"net/http"
	"github.com/gin-gonic/gin"
	"book-store/models"
)

func FindBooks(c *gin.Context){
	var books []models.Books
	models.DB.Find(&books)

	c.JSON(200, gin.H{"data": books})
}