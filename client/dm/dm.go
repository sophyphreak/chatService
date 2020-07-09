package dm

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"../dataForms"
	"../utils"
)

const (
	BaseUrl = "http://localhost:10000"
)

//retrieve list of usernames from server
func GetUserNames() []string {
	//get the response from server
	resp, err := http.Get(BaseUrl + "/username")
	utils.Check(err)
	defer resp.Body.Close()

	users := make([]string, 0)
	//read response and convert to byte
	respBody, err := ioutil.ReadAll(resp.Body)
	utils.Check(err)

	//convert to list
	json.Unmarshal(respBody, &users)

	return users
}

func getInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userName := strings.TrimSpace(scanner.Text())
	return userName
}

func GetConvo(sender, receiver string) dataForms.Conversation {
	page := fmt.Sprintf("/message/direct/%v/%v", sender, receiver)
	resp, err := http.Get(BaseUrl + page)
	utils.Check(err)
	defer resp.Body.Close()

	var convo dataForms.Conversation
	respBody, err := ioutil.ReadAll(resp.Body)
	utils.Check(err)

	json.Unmarshal(respBody, &convo)
	return convo
}

func SendMessage(sender, receiver, body string) {

	dm := dataForms.DirectMessage{Sender: sender, Receiver: receiver, Body: body}
	page := "/message/direct"
	jsonValue, err := json.Marshal(dm)
	utils.Check(err)
	resp, err := http.Post(BaseUrl+page, "application/json", bytes.NewBuffer(jsonValue))
	utils.Check(err)
	resp.Body.Close()
}

func RunDM(sender string) {
	availableUsers := GetUserNames()
	fmt.Println(availableUsers)
ChooseReceiver:
	receiver := getInput("Send message to: ")
	if !utils.StringInSlice(receiver, availableUsers) {
		fmt.Println("That user is not available")
		goto ChooseReceiver
	}
	quit := make(chan byte)
	go listenForDm(sender, receiver, quit)
GetMessage:
	message := getInput("")
	if message == "quit" {
		quit <- 'q'
		return
	}
	if message == "" {
		fmt.Println()
		fmt.Println("Please enter a message!")
		fmt.Println()
		goto GetMessage
	}
	SendMessage(sender, receiver, message)
	goto GetMessage

}
