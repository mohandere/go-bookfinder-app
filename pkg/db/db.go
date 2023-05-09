package db

import (
	"fmt"
	"go-bookfinder-app/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Conn *gorm.DB
)

func Init() *gorm.DB {
	Conn, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Error: Failed to connect database")
	}
	Conn.AutoMigrate(&models.Book{})

	// Create
	Conn.Create(&models.Book{Title: "golang", Description: "This is Desc for go lang book"})
	Conn.Create(&models.Book{Title: "Harry Porter 1", Description: "This is Desc for Harry Porter 1"})
	Conn.Create(&models.Book{Title: "Harry Porter 2", Description: "This is Desc for Harry Porter 2"})
	// db.Find(&book).Debug()
	// fmt.Println("DB book:", book)
	return Conn
}

func PrintDBConn()  {
	fmt.Println("DB conn: ", Conn)
}