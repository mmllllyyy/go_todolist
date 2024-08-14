package main

import (
	"context"
	"todolist/config"
	"todolist/models"
	"todolist/repositories/db"
	"todolist/utils"
)

func main() {
	config.Init()
	utils.LoggerInit()
	db.MySQLInit()

	u := models.User{
		UserName: "mly",
		PassWord: "Dwc258/*-",
		Tasks:    nil,
	}
	ud := db.NewUserDao(context.Background())
	err := ud.CreateUser(&u)
	if err != nil {
		return
	}
	//err := ud.DeleteUser("mly")
	//if err != nil {
	//	return
	//}
}
