package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type channelMessage struct {
	Username string `json:"username"`
	Body     string `json:"body"`
}

func addChannelMessage(w http.ResponseWriter, r *http.Request) {
	var m channelMessage
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.Unmarshal(reqBody, &m)
	message := m // will be replaced with success or failure
	json.NewEncoder(w).Encode(message)
}
