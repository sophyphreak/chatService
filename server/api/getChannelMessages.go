package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func getChannelMessages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	channelName := vars["channelName"]
	// get channel messages
	res := channelName + " slice of messages"
	json.NewEncoder(w).Encode(res)
}
