package main

import (
	"log"
	"os"
	_ "telegram-bot/bot"
	_ "telegram-bot/config"
	_ "telegram-bot/transmission"
	"time"
)

func main() {

	var logFilePath = "./log/telegram_bot.log"

	var logFile *os.File
	if _, err := os.Stat(logFilePath); err != nil {
	} else {
		logFile, _ = os.Create(logFilePath)
		if err != nil {
			log.Fatalln(err)
		}
	}
	//logFile, _ = os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//defer logFile.Close()
	//
	//_ = io.MultiWriter(logFile, os.Stdout)

	log.SetOutput(logFile)
	log.Println("Log File:", logFile)

	log.Println("TEGEGRAM BOT RUNNING!")
	time.Sleep(999999 * time.Hour)

	//service, err := daemon.New("telegram-bot", "텔레그램 봇", daemon.SystemDaemon)
	//if err != nil {
	//	log.Fatal("Error: ", err)
	//}
	//status, err := service.Install()
	//if err != nil {
	//	log.Fatal(status, "\nError: ", err)
	//}
	//fmt.Println("Status:", status)
}
