package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
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
	yfile, err := ioutil.ReadFile("settings.yml")
	if err != nil {
		log.Fatalln(err)
	}

	settings := make(map[string]interface{})
	umsErr := yaml.Unmarshal(yfile, &settings)
	if umsErr != nil {
		log.Fatalln(umsErr)
	}

	CommonFilePath = settings["commonFilePath"].(string)
	SubtitleFilePath = settings["subtitleFilePath"].(string)
	TorrentFilePath = settings["torrentFilePath"].(string)
	SubtitleExts = strings.Split(settings["subtitleExts"].(string), ",")
	TransmissionCommands = strings.Split(settings["transmissionCommands"].(string), ",")

	log.Println("commonFilePath", CommonFilePath)
	log.Println("subtitleFilePath", SubtitleFilePath)
	log.Println("torrentFilePath", TorrentFilePath)
	log.Println("subtitleExts", SubtitleExts)
	log.Println("transmissionCommands", TransmissionCommands)

	Token = os.Getenv("TG_WONA_TOKEN")
	ChatIdStr := os.Getenv("TG_WONA_CHAN_ID")

	chatIdInt, chatIdErr := strconv.Atoi(ChatIdStr)
	if chatIdErr != nil {
		log.Fatalln(chatIdErr)
	}
	ChatId = int64(chatIdInt)
}
