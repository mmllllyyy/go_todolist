package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"todolist/config"
	"todolist/models"
	"todolist/utils"
)

func MySQLInit() {
	cfg := config.DBConf
	dsn := strings.Join([]string{cfg.DbUser, ":", cfg.DbPassWord, "@tcp(", cfg.DBHost, ":", cfg.DbPort, ")/", cfg.DbName, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.Logger.Error("Open database error:" + err.Error())
		return
	} else {
		utils.Logger.Info("Open database done")
	}

	if err = db.AutoMigrate(models.User{}, models.Task{}); err != nil {
		utils.Logger.Error("Failed to migrate models:", err)
	} else {
		utils.Logger.Info("Migrate database done")
	}
}
