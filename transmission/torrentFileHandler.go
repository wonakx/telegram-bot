package transmission

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
	"telegram-bot/config"
	"telegram-bot/util"
)

var TorrentFileChan = make(chan TorrentFile)
var token = config.Token

type TorrentFile struct {
	FileName string
	File     tgbotapi.File
}

func addTorrentFIle() {
	go func() {
		for torrentFile := range TorrentFileChan {
			file := torrentFile.File
			filePath := file.FilePath

			//https://api.telegram.org/file/bot<token>/<file_path>
			urlStrings := []string{"http://api.telegram.org/file/bot", token, "/", filePath}
			url := strings.Join(urlStrings, "")

			torrentFilePath := config.TorrentFilePath + "/" + torrentFile.FileName

			receiveFile := util.GetFileByHttpRequest(url, torrentFilePath)
			log.Println("torrent file received!", receiveFile.Name())

			command, err := util.ExecuteCommand("transmission-remote", "-a", torrentFilePath)
			if err != nil {
				log.Println(err)
			}
			log.Println(command)
		}
	}()
}
