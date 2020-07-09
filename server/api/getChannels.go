package api

import (
	"encoding/json"
	"net/http"

	"../ch"
)

func getChannels(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	channels := ch.GetChannelNames()
	json.NewEncoder(w).Encode(channels)
}
