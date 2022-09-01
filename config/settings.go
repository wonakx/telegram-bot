package config

import (
	"fmt"
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

// 트랜스미션 로그인 정보
var TmPort string
var TmUsername string
var TmPassword string

var LogFilePath = "./log/telegram_bot.log"

func init() {

	SubtitleExts = []string{".srt", ".smi", ".SRT", ".SMI"}
	TransmissionCommands = []string{"list", "add", "del"}

	CommonFilePath = os.Getenv("COMMON_FILE_PATH")
	SubtitleFilePath = os.Getenv("SUB_FILE_PATH")
	TorrentFilePath = os.Getenv("TORRENT_FILE_PATH")

	Token = os.Getenv("TG_TOKEN")
	ChatIdStr := os.Getenv("TG_CHAT_ID")

	chatIdInt, chatIdErr := strconv.Atoi(ChatIdStr)
	if chatIdErr != nil {
		fmt.Print(chatIdErr)
	}
	ChatId = int64(chatIdInt)

	TmPort = os.Getenv("TRANSMISSION_PORT")
	TmUsername = os.Getenv("TRANSMISSION_USERNAME")
	TmPassword = os.Getenv("TRANSMISSION_PASSWORD")

}
