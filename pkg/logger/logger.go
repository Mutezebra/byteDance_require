package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"
)

var LogrusObj *logrus.Logger

func InitLog() {
	if LogrusObj != nil {
		src, _ := setOutFile()
		LogrusObj.Out = src
		log.Println(src)
		return
	}

	logger := logrus.New()
	src, _ := setOutFile()
	logger.Out = src
	fmt.Println(src)

	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	LogrusObj = logger
}

func setOutFile() (*os.File, error) {
	now := time.Now()
	logFilePath := ""

	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "\\logs\\"
	}
	_, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(logFilePath, 0777)
		if err != nil {
			return nil, err
		}
	}
	logFileName := now.Format("2006-01-02") + ".logger"

	fileName := logFilePath + logFileName

	if _, err = os.Stat(fileName); err != nil {
		_, err = os.Create(fileName)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return src, nil
}
