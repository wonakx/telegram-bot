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

	log.Println("TEGEGRAM BOT RUNNING!")

	logFile, err := os.Create("./log/telegram_bot.log")
	if err != nil {
		log.Fatalln(err)
	}
	defer logFile.Close()

	log.Println("Log File:", logFile.Name())

	log.SetOutput(logFile)

	time.Sleep(999999 * time.Hour)
}
