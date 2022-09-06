package transmission

import (
	"github.com/hekmon/transmissionrpc/v2"
	"os"
	"telegram-bot/config"
	"telegram-bot/logwrapper"
)

var log = logwrapper.NewLogger()

var username string
var password string

var Client *transmissionrpc.Client

func init() {

	username = os.Getenv("TRANSMISSION_USERNAME")
	password = os.Getenv("TRANSMISSION_PASSWORD")

	transmissionConfig := transmissionrpc.AdvancedConfig{
		HTTPS:       false,
		Port:        9091,
		HTTPTimeout: 10,
		Debug:       true,
	}

	var err error
	Client, err = transmissionrpc.New("localhost", config.TmUsername, config.TmPassword, &transmissionConfig)
	if err != nil {
		log.Fatalln(err)
	}

	log.Info("[TRANS] username:", username, ", password:", password)

	//Run addTorrentFile
	addTorrentFIle()

	//Run commandHandler
	startCommandHandler()

}
