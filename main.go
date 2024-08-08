package main

import (
	"todolist/config"
	"todolist/repositories/db"
	"todolist/utils"
)

func main() {
	config.Init()
	utils.LoggerInit()
	db.MySQLInit()
}
