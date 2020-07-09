package dm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../utils"
)

const (
	baseUrl = "http://localhost:10000"
)

//retrieve list of usernames from server
func GetUserNames() []string {
	//get the response from server
	resp, err := http.Get(baseUrl + "/username")
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

type Conversation struct {
	User1    string    `json:"user1"`
	User2    string    `json:"user2"`
	Messages []Message `json:"messages"`
}
type Message struct {
	Username string `json:"username"`
	Body     string `json:"body"`
}

// user inputs who they want to talk to
// display all dms between user and other user
func GetConvo(sender, receiver string) Conversation {
	page := fmt.Sprintf("/message/direct/%v/%v", sender, receiver)
	resp, err := http.Get(baseUrl + page)
	utils.Check(err)
	defer resp.Body.Close()

	var convo Conversation
	respBody, err := ioutil.ReadAll(resp.Body)
	utils.Check(err)

	json.Unmarshal(respBody, &convo)
	return convo
}

type directMessage struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Body     string `json:"body"`
}

func SendMessage(sender, receiver, body string) Conversation {

	dm := directMessage{sender, receiver, body}
	page := "/message/direct"
	jsonValue, err := json.Marshal(dm)
	utils.Check(err)
	resp, err := http.Post(baseUrl+page, "application/json", bytes.NewBuffer(jsonValue))
	utils.Check(err)
	defer resp.Body.Close()

	var convo Conversation
	respBody, err := ioutil.ReadAll(resp.Body)
	utils.Check(err)

	json.Unmarshal(respBody, &convo)
	return convo
}

// allow user to input more messages
//		- messages are sent to server
// sends a GET request to /messages/direct/{user}/{otherUser} every .25 seconds
//		- adds any messages that aren't already displayed
// allows user to quit to main menu by typing "quit" or something
