package file

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
	"telegram-bot/config"
	"telegram-bot/util"
)

var SubtitleFileChan = make(chan SubtitleFile)

type SubtitleFile struct {
	FileName string
	File     tgbotapi.File
}

func addSubtitleFIle() {
	go func() {
		for subtitleFile := range SubtitleFileChan {
			file := subtitleFile.File
			filePath := file.FilePath

			//https://api.telegram.org/file/bot<token>/<file_path>
			urlStrings := []string{"http://api.telegram.org/file/bot", config.Token, "/", filePath}
			url := strings.Join(urlStrings, "")

			subtitleFilePath := config.SubtitleFilePath + "/" + subtitleFile.FileName

			receiveFile := util.GetFileByHttpRequest(url, subtitleFilePath)
			log.Info("subtitle file received!", receiveFile.Name())

			FileRespChan <- subtitleFile.FileName + " 파일이 이동 됨."
		}
	}()
}
