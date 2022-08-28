package transmission

import (
	"log"
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
			case "ta":
				log.Print("[" + command.Command + "] receive!\n")
				TransRespChan <- "ta Resp!"
			case "del":
				log.Print("[" + command.Command + "] receive!, parmas: " + strings.Join(command.Parmas, " ") + "\n")
				for _, p := range command.Parmas {
					num, err := strconv.Atoi(p)
					if err != nil {
						log.Println(p, "is not number")
						TransRespChan <- p + " is not number"
					} else {
						log.Println("del", num)
						TransRespChan <- p + " is deleted."
					}
				}
			}

		}
	}(TransChan)
}
