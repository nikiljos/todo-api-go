package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikiljos/todo-api-go/config"
	"github.com/nikiljos/todo-api-go/models"
)

func AddTask(c *gin.Context){

	var task models.Task
	err := c.ShouldBindJSON(&task)
	if err!=nil{
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"success":false,
		})
	}
	dbErr:=config.DB.Create(&task).Error
	if dbErr!=nil{
		fmt.Println(dbErr)
		c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{
			"success":false,
		})
	}
	c.JSON(http.StatusOK,gin.H{
		"success":true,
	})
}

func ListAll(c *gin.Context){
	taskList := &[]models.Task{}
	err:=config.DB.Find(taskList)
	fmt.Println("err",err)
	c.JSON(http.StatusOK,gin.H{
		"success":true,
		"data":taskList,
	})
}