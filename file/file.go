package file

import "telegram-bot/logwrapper"

var log = logwrapper.NewLogger()

func init() {
	addSubtitleFIle()
	addCommonFIle()
}
