package main

import (
	_ "telegram-bot/bot"
	_ "telegram-bot/config"
	_ "telegram-bot/file"
	"telegram-bot/logwrapper"
	_ "telegram-bot/transmission"
	"time"
)

var log = logwrapper.NewLogger()

func main() {

	//var logFilePath = "./log/telegram_bot.log"
	//
	//var logFile *os.File
	//if _, err := os.Stat(logFilePath); err != nil {
	//} else {
	//	logFile, _ = os.Create(logFilePath)
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//}
	//
	//log.SetFormatter(&log.TextFormatter{
	//	FullTimestamp:   true,
	//	TimestampFormat: time.Stamp,
	//})

	//log.SetOutput(os.Stdout)
	//log.SetLevel(log.InfoLevel)

	//log.Info("Log File:", logFile)

	log.Info("TEGEGRAM BOT RUNNING!")
	time.Sleep(999999 * time.Hour)

}
