package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strings"
	"todolist/config"
	"todolist/models"
)

func MySQLInit() {
	cfg := config.DBConf
	dsn := strings.Join([]string{cfg.DbUser, ":", cfg.DbPassWord, "@tcp(", cfg.DBHost, ":", cfg.DbPort, ")/", cfg.DbName, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect mysql:%v", err)
	}

	if err = db.AutoMigrate(models.User{}, models.Task{}); err != nil {
		log.Fatalln("Failed to migrate models: ", err)
	}
}
