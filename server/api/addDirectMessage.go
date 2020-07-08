package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"../cnv"
	"../msg"
)

type directMessage struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Body     string `json:"body"`
}

func addDirectMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var m directMessage
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.Unmarshal(reqBody, &m)

	//add message to golbal var Conversations
	newMessage := msg.Message{Username: m.Sender, Body: m.Body}
	cnv.AddMessage(m.Sender, m.Receiver, newMessage)

	allMessages := cnv.GetMessageList(m.Sender, m.Receiver)
	curConvo := cnv.Conversation{User1: m.Sender, User2: m.Receiver, Messages: allMessages}
	json.NewEncoder(w).Encode(curConvo)
}
