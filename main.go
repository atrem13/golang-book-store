package main

import (
	"book-store/controllers"
	"book-store/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)

	r.Run()
}
