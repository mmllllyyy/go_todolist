package logger

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path"
	"time"
	"todolist/config"
)

var Logger *logrus.Logger

func LoggerInit() {
	logger := logrus.New()

	output := SetLogOutput()
	logger.Out = output

	if config.LogLevel == "Debug" {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	//在输出日志中添加文件名和方法信息
	logger.SetReportCaller(true)

}

func SetLogOutput() *os.File {
	logFilePath := ""
	if root, err := os.Getwd(); err != nil {
		logFilePath = root + "logs"
	}
	_, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		if err := os.Mkdir(logFilePath, 0777); err != nil {
			log.Panicln("Make log dir error:" + err.Error())
		}
	}
	logFileName := time.Now().Format("06-01-02") + ".log"
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if os.IsNotExist(err) {
			if _, err := os.Create(fileName); err != nil {
				log.Panicln("Create log file error:", err.Error())
			}
		}
	}
	file, err := os.Open(logFileName)
	if err != nil {
		log.Panicln("Open log file error:", err.Error())
	}
	return file
}
