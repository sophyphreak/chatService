package api

import (
	"encoding/json"
	"net/http"
)

func getUsernames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := "slice of usernames" // will be replaced
	json.NewEncoder(w).Encode(res)
}
