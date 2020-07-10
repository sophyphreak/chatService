package joinch

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"../utils"
)

// Begin is the entry point to join channel functionality
func Begin(username string) {
	var channels []string
	res, err := http.Get(utils.URL + "/channel")
	if err != nil {
		log.Fatal(err)
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(resBody, &channels)
Choose:
	fmt.Println("List of channels:")
	for _, c := range channels {
		fmt.Println(c)
	}
	fmt.Print("Which channel? ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	channelName := scanner.Text()
	if !utils.StringInSlice(channelName, channels) {
		fmt.Println("That is not an available channel.")
		goto Choose
	}
	quit := make(chan struct{})
	go listen(channelName, quit)
	fmt.Println("If you would like to leave this channel type \"quit\"")
InputMessage:
	fmt.Println()
	scanner.Scan()
	message := strings.TrimSpace(scanner.Text())
	if message == "quit" {
		quit <- struct{}{}
		return
	}

	values := make(map[string]string)
	values["channelName"] = channelName
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
	goto InputMessage
}
