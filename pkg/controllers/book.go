package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-bookfinder-app/pkg/db"
	"go-bookfinder-app/pkg/models"
	"net/http"
)

type BookController struct{}


func (u BookController) Retrieve(c *gin.Context) {
	book := &models.Book{}
	if c.Param("title") != "" {
		fmt.Println("Log: find is -->", db.Conn.Find(&book))
		c.JSON(http.StatusOK, gin.H{"message": "Book found!", "book": db.Conn.First(&book)})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
	return
}
