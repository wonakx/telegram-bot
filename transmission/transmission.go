package transmission

import (
	"log"
	"os"
)

var username string
var password string

func init() {

	username = os.Getenv("TRANSMISSION_USERNAME")
	password = os.Getenv("TRANSMISSION_PASSWORD")

	log.Println("[TRANS] username:", username, ", password:", password)

	//Run addTorrentFile
	addTorrentFIle()

	//Run commandHandler
	startCommandHandler()

}
