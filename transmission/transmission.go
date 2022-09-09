package transmission

import (
	"github.com/hekmon/transmissionrpc/v2"
	"telegram-bot/config"
	"telegram-bot/logwrapper"
	"time"
)

var log = logwrapper.NewLogger()

var Client *transmissionrpc.Client

func init() {

	transmissionConfig := transmissionrpc.AdvancedConfig{
		HTTPS:       true,
		Port:        443,
		HTTPTimeout: 10 * time.Second,
		Debug:       true,
	}

	var err error
	log.Info("userName: " + config.TmUsername + ", password: " + config.TmPassword)
	Client, err = transmissionrpc.New(config.TmHost, config.TmUsername, config.TmPassword, &transmissionConfig)
	if err != nil {
		log.Fatalln(err)
	}

	log.Info("[TRANS] host:", config.TmHost, ", username:", config.TmUsername, ", password:", config.TmPassword)

	//Run addTorrentFile
	addTorrentFIle()

	//Run commandHandler
	startCommandHandler()

}
