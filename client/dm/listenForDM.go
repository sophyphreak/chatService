package dm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"../dataForms"
	"../utils"
)

func listenForDm(user1, user2 string, quit chan byte) {
	messageLen := 0
	page := fmt.Sprintf("/message/direct/%v/%v", user1, user2)
Listen:
	select {
	case <-quit:
		return
	default:
		resp, err := http.Get(utils.URL + page)
		utils.Check(err)
		defer resp.Body.Close()

		var convo dataForms.Conversation
		respBody, err := ioutil.ReadAll(resp.Body)
		utils.Check(err)

		json.Unmarshal(respBody, &convo)
		convoLen := len(convo.Messages)
		if convoLen > messageLen {
			for i := messageLen; i < convoLen; i++ {
				fmt.Printf("%v: %v\n", convo.Messages[i].Username, convo.Messages[i].Body)
			}
			messageLen = convoLen
		}
		time.Sleep(time.Duration(250) * time.Millisecond)
		goto Listen
	}
}
