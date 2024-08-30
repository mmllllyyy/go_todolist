package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	UserName string `gorm:"unique"`
	PassWord string `gorm:"not null"`
}

type Task struct {
	gorm.Model
	UserID   uint
	TaskName string    `gorm:"not null"`
	Status   bool      `gorm:"default:0"`
	Content  string    `gorm:"type:longtext"`
	StartAt  time.Time `time_format:"2006-01-02 15:04"`
	EndAt    time.Time `time_format:"2006-01-02 15:04"`
	User     User      `gorm:"ForeignKey:UserID"`
}
