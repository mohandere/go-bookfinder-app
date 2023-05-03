package db

import (
	"go-playground/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Error: Failed to connect database")
	}

	db.AutoMigrate(&models.Book{})

	// Create
	db.Create(&models.Book{Title: "Harry Porter 1", Description: "This is Desc for Harry Porter 1"})
	db.Create(&models.Book{Title: "Harry Porter 2", Description: "This is Desc for Harry Porter 2"})
	return db
}
func GetDB() *gorm.DB {
	return db
}
