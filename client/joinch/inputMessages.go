package joinch

import (
	"bufio"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"../utils"
)

func inputMessages(currentChannel, username string) {
	scanner := bufio.NewScanner(os.Stdin)
Top:
	scanner.Scan()
	message := strings.TrimSpace(scanner.Text())
	if message == "quit" {
		return
	}
	sendMessage(currentChannel, username, message)
	goto Top
}

func sendMessage(currentChannel, username, message string) {
	values := make(map[string]string)
	values["channelName"] = currentChannel
	values["username"] = username
	values["body"] = message
	jsonValue, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}
	_, err = http.Post(utils.URL+"/message/channel", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Fatal(err)
	}
}
