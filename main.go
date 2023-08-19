package main

import (
	config "ujiKeterampilan/config"
	"ujiKeterampilan/controllers/bookController"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	router := gin.Default()

	api := router.Group("api/v1")
	api.GET("/books", bookController.Index)
	api.GET("/book/:id", bookController.Show)
	api.POST("/book", bookController.Create)
	// api.PUT("/book/:id", bookController.Update)
	api.DELETE("/book/:id", bookController.Delete)

	router.Run()
}
