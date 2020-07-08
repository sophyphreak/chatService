package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func getDirectMessages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user1 := vars["user1"]
	user2 := vars["user2"]
	if user2 < user1 {
		user1, user2 = user2, user1
	}

	// add get direct messages function

	res := make(map[string]string)
	res["user1"] = user1
	res["user2"] = user2
	res["messages"] = "slice of messages"
	json.NewEncoder(w).Encode(res)
}
