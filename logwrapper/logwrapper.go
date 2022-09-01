package logwrapper

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
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

	var logFilePath = "./log/telegram_bot.log"

	var logFile *os.File
	if _, err := os.Stat(logFilePath); err != nil {
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
		FullTimestamp:   true,
		TimestampFormat: time.Stamp,
	}
	return standardLogger
}
