package models 

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql@v1.9.16"
)

var DB *gorm.DB

func ConnectDatabase(){
	database, errr := gorm.Open("mysql", "trygolang")

	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(&Book{})

	DB = database
}