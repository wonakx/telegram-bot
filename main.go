package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	_ "telegram-bot/bot"
	_ "telegram-bot/config"
	_ "telegram-bot/transmission"
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

	// load command line arguments
	name := flag.String("name", "world", "name to print")
	flag.Parse()

	log.Printf("Starting sleepservice for %s", *name)

	// setup signal catching
	sigs := make(chan os.Signal, 1)

	// catch all signals since not explicitly listing
	signal.Notify(sigs)
	//signal.Notify(sigs,syscall.SIGQUIT)

	// method invoked upon seeing signal
	s := <-sigs
	log.Printf("RECEIVED SIGNAL: %s", s)
	AppCleanup()
	os.Exit(1)

}

func AppCleanup() {
	log.Println("CLEANUP APP BEFORE EXIT!!!")
}
