package transmission

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

var TransChan = make(chan TransCommand)

type TransCommand struct {
	Command string
	Parmas  []string
}

func (tc *TransCommand) listTorrents() string {
	log.Info("[" + tc.Command + "] receive!\n")

	torrents, err := Client.TorrentGetAll(context.TODO())
	if err != nil {
		log.Error(err)
	}

	var listStat []string
	for _, torrent := range torrents {
		id := strconv.FormatInt(*torrent.ID, 10)
		name := *torrent.Name
		status := torrent.Status.String()
		download := fmt.Sprintf("%.1f %s", *torrent.PercentDone*100, "%")

		stat := strings.Join([]string{"[", id, "]", name, status, "[", download, "]"}, " ")
		listStat = append(listStat, stat)
	}

	return strings.Join(listStat, "\n")
}

func startCommandHandler() {
	commandHandler()
}

func commandHandler() {
	go func(transChan <-chan TransCommand) {
		for command := range transChan {
			switch command.Command {
			case "list", "li":
				TransRespChan <- command.listTorrents()
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
