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
)

// display list of channels
// user inputs what channel they want to join
// display all messages in channel
// allow user to input more messages into channel
//		- messages are sent to channel on server
// sends a GET request to /messages/channel/{channelName} every .25 seconds
//		- adds any messages that aren't already displayed
// allows user to quit to main menu by typing "quit" or something

// Begin is the entry point to join channel functionality
func Begin(username string) {
	var channels []string
	res, err := http.Get("http://localhost:10000/channel")
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
	if !stringInSlice(channelName, channels) {
		fmt.Println("That is not an available channel.")
		goto Choose
	}
	quit := make(chan struct{})
	go listen(channelName, quit)
InputMessage:
	fmt.Println("If you would like to leave this channel type \"quit\"")
	fmt.Print("What message would you like to send? ")
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
	_, err = http.Post("http://localhost:10000/message/channel", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Fatal(err)
	}
	goto InputMessage
}
