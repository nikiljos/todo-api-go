package config

import (
	"fmt"
	"os"

	"github.com/nikiljos/todo-api-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func OpenDb() bool{
	dsn:=os.Getenv("PG_DSN")

	db,err := gorm.Open(postgres.Open(dsn))
	if err!=nil {
		fmt.Println("DB Connect Error:",err)
		return false
	}
		
	DB=db
	return true
}

func AutoMigrate(){
	DB.AutoMigrate(&models.Task{})
}