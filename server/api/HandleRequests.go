package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// HandleRequests handles API requests
func HandleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/message/channel/{channelName}", getChannelMessages).Methods("GET")
	router.HandleFunc("/message/direct/{user1}/{user2}", getDirectMessages).Methods("GET")
	router.HandleFunc("/channel", getChannels).Methods("GET")
	router.HandleFunc("/username", getUsernames).Methods("GET")
	router.HandleFunc("/message/channel", addChannelMessage).Methods("POST")
	router.HandleFunc("/message/direct/{user1}/{user2}", addDirectMessage).Methods("POST")
	router.HandleFunc("/channel", addChannel).Methods("POST")
	router.HandleFunc("/username", addUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", router))
}
