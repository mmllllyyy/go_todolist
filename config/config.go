package config

import (
	"gopkg.in/ini.v1"
	"log"
)

var LogLevel string

func LoadLogConf(file *ini.File) {
	LogLevel = file.Section("LogLevel").Key("Level").String()
}

type DBConfType struct {
	DBHost     string `ini:"DBHost"`
	DbPort     string `ini:"DBPort"`
	DbPassWord string `ini:"DBPassWord"`
	DbUser     string `ini:"DBUser"`
	DbName     string `ini:"DBName"`
}

var DBConf DBConfType

func LoadDBConf(file *ini.File) {
	err := file.Section("MySQL").MapTo(&DBConf)
	if err != nil {
		log.Fatalln("Failed to map struct:" + err.Error())
		return
	}
}

func Init() {
	file, err := ini.Load("./config/config.init")
	if err != nil {
		log.Fatalln("Failed to load config.init:", err)
	}
	LoadLogConf(file)
	LoadDBConf(file)
}
