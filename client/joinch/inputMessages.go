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
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		message := strings.TrimSpace(scanner.Text())
		if message == "quit" {
			return
		}
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
}
