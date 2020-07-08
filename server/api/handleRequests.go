package api

import (
	"github.com/gorilla/mux"
)

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/message/channel", getChannelMessages).Methods("GET")
	router.HandleFunc("/message/direct", getDirectMessages).Methods("GET")
	router.HandleFunc("/channel", getChannels).Methods("GET")
	router.HandleFunc("/username", getUsernames).Methods("GET")
	router.HandleFunc("/message/channel", addChannelMessage).Methods("POST")
	router.HandleFunc("/message/direct", addDirectMessage).Methods("POST")
	router.HandleFunc("/channel", addChannel).Methods("POST")
	router.HandleFunc("/username", addUser).Methods("POST")
}
