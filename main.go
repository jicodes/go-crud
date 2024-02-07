package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jicodes/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	router := gin.Default()
	
	Routes(router)

	router.Run()
}