package api

import (
	"encoding/json"
	"net/http"

	"../usr"
)

func getUsernames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := usr.GetUserList()
	json.NewEncoder(w).Encode(res)
}
