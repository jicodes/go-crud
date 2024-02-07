package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jicodes/go-crud/controllers"
)

func Routes(router *gin.Engine) {

	router.POST("/posts", controllers.CreatePost)
	router.GET("/posts", controllers.GetPosts)
	router.GET("/posts/:id", controllers.GetPost)
	router.PUT("/posts/:id", controllers.UpdatePost)
	router.DELETE("/posts/:id", controllers.DeletePost)
}