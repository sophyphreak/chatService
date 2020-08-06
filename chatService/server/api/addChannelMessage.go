package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"../ch"
	"../msg"
)

type channelMessage struct {
	ChannelName string `json:"channelName"`
	Username    string `json:"username"`
	Body        string `json:"body"`
}

func addChannelMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var m channelMessage
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.Unmarshal(reqBody, &m)

	message := msg.CreateMessage(m.Username, m.Body)
	curChannel := ch.AddMessage(message, m.ChannelName)
	json.NewEncoder(w).Encode(curChannel)
}
