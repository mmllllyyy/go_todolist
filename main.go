package main

import (
	"todolist/config"
	"todolist/repositories/db"
	"todolist/routes"
	"todolist/utils"
)

func main() {
	config.Init()
	utils.LoggerInit()
	db.MySQLInit()

	r := routes.NewRouter()
	_ = r.Run(":3000")
	//u := models.User{
	//	UserName: "mly",
	//	PassWord: "Dwc258/*-",
	//	Tasks:    nil,
	//}
	//ud := db.NewUserDao(context.Background())
	//err := ud.CreateUser(&u)
	//if err != nil {
	//	return
	//}
	//err := ud.DeleteUser("mly")
	//if err != nil {
	//	return
	//}
}
