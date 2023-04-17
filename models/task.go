package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	ID        uint `gorm:"primary_key"`
	Title     string `json:"title" binding:"required"`
	Completed bool
}