package bot

import (
	"bytes"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os/exec"
	"strings"
	"telegram-bot/config"
	"telegram-bot/file"
	"telegram-bot/logwrapper"
	"telegram-bot/transmission"
	"telegram-bot/util"
	"time"
)

var log = logwrapper.NewLogger()

var Bot *tgbotapi.BotAPI

var transChan = transmission.TransChan
var torrFileChan = transmission.TorrentFileChan
var subFileChan = file.SubtitleFileChan
var CommonFileChan = file.CommonFileChan

func init() {

	Bot, _ = tgbotapi.NewBotAPI(config.Token)

	log.Printf("Authorized on account %s", Bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := Bot.GetUpdatesChan(u)

	go func() {
		for update := range updates {
			if update.Message != nil {
				log.Infof("[%s] %s", update.Message.From.UserName, update.Message.Text)

				go commandControl(update.Message.Text)

				go fileControl(update.Message.Document)

				go commandResp()
			}
		}
	}()

	checkCurrentList()
}

func checkCurrentList() {
	for {
		cmd := exec.Command("transmission-remote", config.TmPort, "--auth", config.TmUsername+":"+config.TmPassword, "-l")
		var outb, errb bytes.Buffer
		cmd.Stdout = &outb
		cmd.Stderr = &errb

		if err := cmd.Run(); err != nil {
			log.Error(err)
		}

		torrentList := outb.String()
		split := strings.Split(torrentList, "\n")

		var endIndex int
		if (len(split) - 2) >= 1 {
			endIndex = len(split) - 2
		} else {
			endIndex = 1
		}

		for _, row := range split[1:endIndex] {
			fields := strings.Fields(row)
			id := fields[0]
			progress := fields[1]
			log.Infoln("ID:", id, "progress:", progress)
			for idx, field := range fields {
				log.Info("idx:", idx, "field:", field)
			}
			log.Infoln()
		}

		time.Sleep(10 * time.Second)
	}
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
		log.Info("Send Command", command, "to TransChan")
	}
}

func fileControl(document *tgbotapi.Document) {
	if document != nil {
		fileConfig := tgbotapi.FileConfig{FileID: document.FileID}
		recvFile, fileErr := Bot.GetFile(fileConfig)

		if fileErr != nil {
			log.Error(fileErr)
		}

		if strings.HasSuffix(document.FileName, ".torrent") {
			torrentFile := transmission.TorrentFile{
				FileName: document.FileName,
				File:     recvFile,
			}
			torrFileChan <- torrentFile

		} else if util.ContainsEndWith(config.SubtitleExts, document.FileName) {
			log.Info("It is not torrent recvFile.", document.FileName)
			subtitleFile := file.SubtitleFile{
				FileName: document.FileName,
				File:     recvFile,
			}
			subFileChan <- subtitleFile

		} else {
			log.Info("It is normal recvFile.", document.FileName)
			commonFile := file.CommonFile{
				FileName: document.FileName,
				File:     recvFile,
			}
			CommonFileChan <- commonFile
		}
	}
}

var transRespChan = transmission.TransRespChan

var fileRespChan = file.FileRespChan

func commandResp() {
	for {
		select {
		case transResp := <-transRespChan:
			_, err := Bot.Send(tgbotapi.NewMessage(config.ChatId, transResp))
			if err != nil {
				log.Info("Send Error!", err)
			}
		case fileResp := <-fileRespChan:
			_, err := Bot.Send(tgbotapi.NewMessage(config.ChatId, fileResp))
			if err != nil {
				log.Info("Send Error!", err)
			}
		}
	}
}
