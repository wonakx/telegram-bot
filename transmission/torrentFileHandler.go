package transmission

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
	"telegram-bot/config"
	"telegram-bot/util"
)

var token = config.Token

var TorrentFileChan = make(chan TorrentFile)

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

			torrentFilePath := config.CommonFilePath + "/" + torrentFile.FileName

			receiveFile := util.GetFileByHttpRequest(url, torrentFilePath)
			log.Info("torrent file received!", receiveFile.Name())

			var respMsg string
			torrent, taddErr := Client.TorrentAddFile(context.TODO(), torrentFilePath)
			if taddErr != nil {
				respMsg = torrentFile.FileName + " 파일 추가 실패. " + taddErr.Error()
				log.Error(taddErr)
			} else {
				respMsg = "[" + strconv.FormatInt(*torrent.ID, 10) + "] " + *torrent.Name + " 파일이 성공적으로 추가 됨."
			}
			TransRespChan <- respMsg
		}
	}()
}
