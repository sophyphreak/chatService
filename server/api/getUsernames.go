package api

import (
	"encoding/json"
	"net/http"
)

func getUsernames(w http.ResponseWriter, r *http.Request) {
	res := "slice of usernames" // will be replaced
	json.NewEncoder(w).Encode(res)
}
