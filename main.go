package main

import (
	config "ujiKeterampilan/config"
	"ujiKeterampilan/controllers/bookController"
	"ujiKeterampilan/controllers/loanController"
	"ujiKeterampilan/controllers/memberController"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	router := gin.Default()

	api := router.Group("api/v1")
	api.GET("/books", bookController.Index)
	api.GET("/book/:id", bookController.Show)
	api.POST("/book", bookController.Create)
	api.PUT("/book/:id", bookController.Update)
	api.DELETE("/book/:id", bookController.Delete)

	api.GET("/members", memberController.Index)
	api.GET("/member/:id", memberController.Show)
	api.POST("/member", memberController.Create)
	api.PUT("/member/:id", memberController.Update)
	api.DELETE("/member/:id", memberController.Delete)

	api.GET("/loans", loanController.Index)
	api.GET("/loan/:id", loanController.Show)
	api.POST("/loan", loanController.Create)
	api.DELETE("/loan/:id", loanController.Delete)

	router.Run()
}
