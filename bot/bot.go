package bot

import (
	"bytes"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os/exec"
	"strings"
	"telegram-bot/config"
	"telegram-bot/file"
	"telegram-bot/transmission"
	"telegram-bot/util"
	"time"
)

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
				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

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
			log.Println(err)
		}

		log.Println("Cmd Result:\n", outb.String())

		torrentList := outb.String()
		split := strings.Split(torrentList, "\n")
		log.Println(strings.Join(split, " "))
		log.Println(len(split))
		for k, row := range split[1 : len(split)-1] {
			log.Println("rowidx:", k, "value:", row, " ")
			fields := strings.Fields(row)
			for idx, field := range fields {
				log.Println("idx:", idx, "field:", field)
			}
			//id := fields[0]
			//done := fields[1]
			//status := fields[6]
			//log.Println("id:", id, "done:", done, "status:", status)
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

var transRespChan = transmission.TransRespChan
var torrentRespChan = transmission.TorrentRespChan

var subtitleRespChan = file.SubtitleFileRespChan
var commonRespChan = file.CommonFileRespChan

func commandResp() {
	for {
		select {
		case transResp := <-transRespChan:
			_, err := Bot.Send(tgbotapi.NewMessage(config.ChatId, transResp))
			if err != nil {
				log.Println("Send Error!", err)
			}
		case torrentFileResp := <-torrentRespChan:
			_, err := Bot.Send(tgbotapi.NewMessage(config.ChatId, torrentFileResp))
			if err != nil {
				log.Println("Send Error!", err)
			}
		case subtitleResp := <-subtitleRespChan:
			_, err := Bot.Send(tgbotapi.NewMessage(config.ChatId, subtitleResp))
			if err != nil {
				log.Println("Send Error!", err)
			}
		case commonFileResp := <-commonRespChan:
			_, err := Bot.Send(tgbotapi.NewMessage(config.ChatId, commonFileResp))
			if err != nil {
				log.Println("Send Error!", err)
			}
		}
	}
}
