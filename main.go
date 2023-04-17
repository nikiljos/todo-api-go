package main

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nikiljos/todo-api-go/config"
	"github.com/nikiljos/todo-api-go/controllers"
)

func main() {

	godotenv.Load()
	config.OpenDb()
	config.AutoMigrate()

	router := gin.Default()

	router.GET("/ping",controllers.Ping)
	router.GET("/task",controllers.ListAll)
	router.POST("/task",controllers.AddTask)

	router.Run(":3000")
}