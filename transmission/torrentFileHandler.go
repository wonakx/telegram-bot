package transmission

import (
	"bytes"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os/exec"
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

			torrentFilePath := config.CommonFilePath + "/" + torrentFile.FileName

			receiveFile := util.GetFileByHttpRequest(url, torrentFilePath)
			log.Info("torrent file received!", receiveFile.Name())

			//command, err := util.ExecuteCommand("transmission-remote", config.TmPort, "--auth", config.TmUsername+":"+config.TmPassword, "-a", torrentFilePath)
			//if err != nil {
			//	log.Error(err)
			//}
			//log.Info(command)

			cmd := exec.Command("transmission-remote", config.TmPort, "--auth", config.TmUsername+":"+config.TmPassword, "-a", torrentFilePath)
			var outb, errb bytes.Buffer
			cmd.Stdout = &outb
			cmd.Stderr = &errb
			err := cmd.Run()
			if err != nil {
				log.Error(err)
			}
			result := outb.String()
			fields := strings.Fields(result)
			success := fields[len(fields)-1]
			fold := strings.EqualFold(success, "\"success\"")
			if fold {
				TorrentRespChan <- torrentFile.FileName + " 파일이 성공적으로 추가 됨."
			} else {
				TorrentRespChan <- torrentFile.FileName + " 파일이 추가 중 오류 발생."
			}
		}
	}()
}
