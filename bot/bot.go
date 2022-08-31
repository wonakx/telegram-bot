package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
	"telegram-bot/config"
	"telegram-bot/file"
	"telegram-bot/transmission"
	"telegram-bot/util"
)

var Bot *tgbotapi.BotAPI

var transChan = transmission.TransChan
var torrFileChan = transmission.TorrentFileChan
var subFileChan = file.SubtitleFileChan
var CommonFileChan = file.CommonFileChan

var transRespChan = transmission.TransRespChan
var torrentRespChan = transmission.TorrentRespChan

func init() {

	Bot, _ = tgbotapi.NewBotAPI(config.Token)

	log.Printf("Authorized on account %s", Bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := Bot.GetUpdatesChan(u)

	go func() {
		for update := range updates {
			if update.Message != nil {
				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

				go commandControl(update.Message.Text)

				go fileControl(update.Message.Document)

				go commandResp()
			}
		}
	}()
}

func commandControl(command string) {
	switch util.ContainsStartWith(config.TransmissionCommands, command) {
	case true:
		commandFactors := strings.Split(command, " ")
		transCommand := transmission.TransCommand{
			Command: commandFactors[0],
			Parmas:  commandFactors[1:],
		}
		transChan <- transCommand
		log.Println("Send Command", command, "to TransChan")
	}
}

func fileControl(document *tgbotapi.Document) {
	if document != nil {
		fileConfig := tgbotapi.FileConfig{FileID: document.FileID}
		recvFile, fileErr := Bot.GetFile(fileConfig)

		if fileErr != nil {
			log.Fatalln(fileErr)
		}

		if strings.HasSuffix(document.FileName, ".torrent") {
			torrentFile := transmission.TorrentFile{
				FileName: document.FileName,
				File:     recvFile,
			}
			torrFileChan <- torrentFile

		} else if util.ContainsEndWith(config.SubtitleExts, document.FileName) {
			log.Println("It is not torrent recvFile.", document.FileName)
			subtitleFile := file.SubtitleFile{
				FileName: document.FileName,
				File:     recvFile,
			}
			subFileChan <- subtitleFile

		} else {
			log.Println("It is normal recvFile.", document.FileName)
			commonFile := file.CommonFile{
				FileName: document.FileName,
				File:     recvFile,
			}
			CommonFileChan <- commonFile
		}
	}
}

func commandResp() {
	for {
		select {
		case transResp := <-transRespChan:
			send, err := Bot.Send(tgbotapi.NewMessage(config.ChatId, transResp))
			if err != nil {
				log.Println("Send Error!", err)
			}
			log.Println(send.Text, "send success!")
		case torrentFileResp := <-torrentRespChan:
			send, err := Bot.Send(tgbotapi.NewMessage(config.ChatId, torrentFileResp))
			if err != nil {
				log.Println("Send Error!", err)
			}
			log.Println(send.Text, "send success!")
		}
	}
}
