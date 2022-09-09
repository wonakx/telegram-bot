package file

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
	"telegram-bot/config"
	"telegram-bot/util"
)

var CommonFileChan = make(chan CommonFile)

type CommonFile struct {
	FileName string
	File     tgbotapi.File
}

func addCommonFIle() {
	go func() {
		for commonFile := range CommonFileChan {
			file := commonFile.File
			filePath := file.FilePath

			//https://api.telegram.org/file/bot<token>/<file_path>
			urlStrings := []string{"http://api.telegram.org/file/bot", config.Token, "/", filePath}
			url := strings.Join(urlStrings, "")

			commonFilePath := config.CommonFilePath + "/" + commonFile.FileName

			receiveFile := util.GetFileByHttpRequest(url, commonFilePath)
			log.Info("common file received!", receiveFile.Name())

			FileRespChan <- commonFile.FileName + " 파일이 이동 됨."
		}
	}()
}
