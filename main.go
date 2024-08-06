package main

import (
	"todolist/config"
	"todolist/repositories/db"
)

func main() {
	config.Init()
	db.MySQLInit()
}
