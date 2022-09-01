package transmission

import (
	"os"
	"telegram-bot/logwrapper"
)

var log = logwrapper.NewLogger()

var username string
var password string

func init() {

	username = os.Getenv("TRANSMISSION_USERNAME")
	password = os.Getenv("TRANSMISSION_PASSWORD")

	log.Info("[TRANS] username:", username, ", password:", password)

	//Run addTorrentFile
	addTorrentFIle()

	//Run commandHandler
	startCommandHandler()

}
