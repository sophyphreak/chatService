package api

import (
	"encoding/json"
	"net/http"
)

func getChannels(w http.ResponseWriter, r *http.Request) {
	res := "slice of channel names" // will be replaced
	json.NewEncoder(w).Encode(res)
}
