package config

import (
	"log"
	"os"
	"strconv"
)

var ChatId int64
var Token string

// 파일 디렉토리 경로
var CommonFilePath string
var SubtitleFilePath string
var TorrentFilePath string
var TransmissionCommands []string

// 자막 파일 확장자 리스트
var SubtitleExts []string

func init() {

	SubtitleExts = []string{".srt", ".smi", ".SRT", ".SMI"}
	TransmissionCommands = []string{"list", "add", "del"}

	log.Println("subtitleExts", SubtitleExts)
	log.Println("transmissionCommands", TransmissionCommands)

	CommonFilePath = os.Getenv("COMMON_FILE_PATH")
	SubtitleFilePath = os.Getenv("SUB_FILE_PATH")
	TorrentFilePath = os.Getenv("TORRENT_FILE_PATH")

	log.Println("commonFilePath", CommonFilePath)
	log.Println("subtitleFilePath", SubtitleFilePath)
	log.Println("torrentFilePath", TorrentFilePath)

	Token = os.Getenv("TG_TOKEN")
	ChatIdStr := os.Getenv("TG_CHAT_ID")

	log.Println("TG_TOKEN:", Token)
	log.Println("TG_CHAI_ID:", ChatIdStr)

	chatIdInt, chatIdErr := strconv.Atoi(ChatIdStr)
	if chatIdErr != nil {
		log.Fatalln(chatIdErr)
	}
	ChatId = int64(chatIdInt)
}
