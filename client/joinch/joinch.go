package joinch

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
func Begin() {
	var channels []string
	res, err := http.Get("http://localhost:10000/channel")
	if err != nil {
		log.Fatal(err)
	}
	reqBody, err := ioutil.ReadAll(res.Body)
	json.Unmarshal(reqBody, &channels)
	fmt.Println("List of channels:")
	for _, c := range channels {
		fmt.Println(c)
	}
	fmt.Print("Which channel? ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	channelName := scanner.Text()
	fmt.Println(channelName)

}
