package main

import (
	"github.com/jicodes/go-crud/initializers"
	"github.com/jicodes/go-crud/models"
)

func init () {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}