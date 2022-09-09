package logwrapper

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"telegram-bot/config"
	"time"
)

type Event struct {
	id      int
	message string
}

type StandardLogger struct {
	*logrus.Logger
}

func NewLogger() *StandardLogger {
	var baseLogger = logrus.New()

	var logFilePath = config.LogFilePath

	var logFile *os.File
	if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
		logFile, _ = os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	} else {
		logFile, _ = os.Create(logFilePath)
		if err != nil {
			log.Fatalln(err)
		}
	}

	baseLogger.Out = logFile

	var standardLogger = &StandardLogger{baseLogger}
	standardLogger.Formatter = &logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	}
	return standardLogger
}
