package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getChannelMessages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user1 := vars["user1"]
	user2 := vars["user2"]
	fmt.Println(user1)
	fmt.Println(user2)

	// gameIdInt, _ := strconv.Atoi(gameId)
	// foundGame, err := game.GetGame(gameIdInt)
	// game.Check(err)
	// json.NewEncoder(w).Encode(foundGame)

	// receives
	//  - channel name

	// returns
	//  - slice of messages
}
