package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"../cnv"
	"../msg"
)

type directMessageList struct {
	User1    string        `json:"user1"`
	User2    string        `json:"user2"`
	Messages []msg.Message `json:"messages"`
}

func getDirectMessages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user1 := vars["user1"]
	user2 := vars["user2"]
	if user2 < user1 {
		user1, user2 = user2, user1
	}

	var m directMessageList
	m.User1 = user1
	m.User2 = user2
	m.Messages = cnv.GetMessageList(user1, user2)
	json.NewEncoder(w).Encode(m)
}
