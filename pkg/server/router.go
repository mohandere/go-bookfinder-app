package server

import (
	"github.com/gin-gonic/gin"
	"go-bookfinder-app/pkg/controllers"
	"net/http"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.LoadHTMLGlob("web/templates/*")

	v1 := router.Group("v1")
	{
		bookGroup := v1.Group("book")
		{
			book := new(controllers.BookController)
			bookGroup.GET("/:title", book.Retrieve)
		}
	}

	// router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Hello world!",
		})
	})

	return router

}
