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
	type ResData struct{
		ID uint `json:"id"`
		Title string `json:"title"`
		Completed bool `json:"completed"`
	}
	var taskList []ResData	
	err:=config.DB.Model(&models.Task{}).Find(&taskList).Error
	if err!=nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{
			"success":false,
		})
	}
	c.JSON(http.StatusOK,gin.H{
		"success":true,
		"data":taskList,
	})
}