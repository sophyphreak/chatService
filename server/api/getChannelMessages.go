package api

import (
	"encoding/json"
	"net/http"

	"../ch"
	"../msg"
	"github.com/gorilla/mux"
)

type response struct {
	Messages []msg.Message `json:"messages"`
}

func getChannelMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	channelName := vars["channelName"]
	msgs := ch.GetMessages(channelName)
	var messages response
	messages.Messages = msgs
	json.NewEncoder(w).Encode(messages)
}
