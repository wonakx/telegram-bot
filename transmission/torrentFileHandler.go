package transmission

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
	"telegram-bot/config"
	"telegram-bot/util"
)

var token = config.Token

var TorrentFileChan = make(chan TorrentFile)
var TorrentRespChan = make(chan string)

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
			log.Info("torrent file received!", receiveFile.Name())

			command, err := util.ExecuteCommand("transmission-remote", config.TmPort, "--auth", config.TmUsername+":"+config.TmPassword, "-a", torrentFilePath)
			if err != nil {
				log.Error(err)
			}
			log.Info(command)

			TorrentRespChan <- torrentFile.FileName + " watch 디렉토리로 파일이 이동 됨."
		}
	}()
}
