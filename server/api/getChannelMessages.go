package api

import (
	"encoding/json"
	"net/http"

	"../ch"
	"github.com/gorilla/mux"
)

func getChannelMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	channelName := vars["channelName"]
	messages := ch.GetMessages(channelName)
	json.NewEncoder(w).Encode(messages)
}
