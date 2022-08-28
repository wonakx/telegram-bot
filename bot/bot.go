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

func init() {

	var token = config.Token
	var chatId = config.ChatId

	log.Println("[TELEGRAM INFO] tgToken:", token, "ChatId:", chatId)

	Bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	Bot.Debug = false

	log.Printf("Authorized on account %s", Bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := Bot.GetUpdatesChan(u)

	transChan := transmission.TransChan
	torrFileChan := transmission.TorrentFileChan
	subFileChan := file.SubtitleFileChan

	transRespChan := transmission.TransRespChan
	go func() {
		for update := range updates {
			if update.Message != nil {
				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

				go func(command string) {
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
				}(update.Message.Text)

				go func(document *tgbotapi.Document) {

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
						}

					} else {
						log.Println("Not File!")
					}
				}(update.Message.Document)

				go func() {
					for {
						select {
						case transResp := <-transRespChan:
							Bot.Send(tgbotapi.NewMessage(chatId, transResp))
						}
					}
				}()
			}
		}
	}()
}
