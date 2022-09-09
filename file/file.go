package file

import "telegram-bot/logwrapper"

var log = logwrapper.NewLogger()

var FileRespChan = make(chan string)

func init() {
	addSubtitleFIle()
	addCommonFIle()
}
