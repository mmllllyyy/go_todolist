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
}
