package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context){
	fmt.Println(c.Request.Header)
	c.JSON(200,gin.H{
		"success":true,
		"message":"pong",
	})
}