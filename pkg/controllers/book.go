package controllers

import (
	"github.com/gin-gonic/gin"
	"go-playground/pkg/db"
	"go-playground/pkg/models"
	"net/http"
)

type BookController struct{}

var bookModel models.Book

func (u BookController) Retrieve(c *gin.Context) {
	if c.Param("id") != "" {
		db := db.GetDB()
		c.JSON(http.StatusOK, gin.H{"message": "Book found!", "book": db.First(&bookModel)})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
	return
}
