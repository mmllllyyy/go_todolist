package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	UserName string `gorm:"unique"`
	PassWord string
	Tasks    []Task
}

type Task struct {
	gorm.Model
	UserID   int
	TaskName string `gorm:"not null"`
	Status   bool   `gorm:"default:0"`
	Content  string `gorm:"type:longtext"`
	StartAt  time.Time
	EndAt    time.Time
}
