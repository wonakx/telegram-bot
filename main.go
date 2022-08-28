package main

import (
	"log"
	_ "telegram-bot/bot"
	_ "telegram-bot/config"
	_ "telegram-bot/transmission"
	"time"
)

func main() {

	log.Println("TEGEGRAM BOT RUNNING!")

	time.Sleep(999999 * time.Hour)
}
