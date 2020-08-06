package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"../ch"
)

type channelName struct {
	ChannelName string `json:"name"`
}

func addChannel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var name channelName
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.Unmarshal(reqBody, &name)
	curChannel := ch.AddChannel(name.ChannelName)
	json.NewEncoder(w).Encode(curChannel)
}
