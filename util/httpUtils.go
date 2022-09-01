package util

import (
	"io/ioutil"
	"net/http"
	"os"
	"telegram-bot/logwrapper"
)

var log = logwrapper.NewLogger()

func GetFileByHttpRequest(url string, destFilePath string) os.File {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{}
	resp, respErr := client.Do(req)
	if respErr != nil {
		log.Fatalln(respErr)
	}

	bytes, err := ioutil.ReadAll(resp.Body)

	err = ioutil.WriteFile(destFilePath, bytes, 0644)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}

	openFile, err := os.Open(destFilePath)
	if err != nil {
		log.Fatalln(err)
	}

	return *openFile
}
