package transmission

import (
	"strconv"
	"strings"
)

var TransChan = make(chan TransCommand)

var TransRespChan = make(chan string)

type TransCommand struct {
	Command string
	Parmas  []string
}

func startCommandHandler() {
	commandHandler()
}

func commandHandler() {
	go func(transChan <-chan TransCommand) {
		for command := range transChan {
			switch command.Command {
			case "list":
				log.Info("[" + command.Command + "] receive!\n")
				TransRespChan <- "list Resp!"
			case "del":
				log.Info("[" + command.Command + "] receive!, parmas: " + strings.Join(command.Parmas, " ") + "\n")
				for _, p := range command.Parmas {
					num, err := strconv.Atoi(p)
					if err != nil {
						log.Error(p, "is not number")
						TransRespChan <- p + " is not number"
					} else {
						log.Error("del", num)
						TransRespChan <- p + " is deleted."
					}
				}
			}

		}
	}(TransChan)
}
