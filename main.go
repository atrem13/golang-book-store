package main


import (
	"net/http"
	"github.com/gin-gonic/gin"
	"book-store/models"
	"book-store/controllers"
)

func main(){
	r := gin.Default()

	models.ConnectDatabase()
	
	r.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.GET("/books", controllers.FindBooks)

	r.Run()
}