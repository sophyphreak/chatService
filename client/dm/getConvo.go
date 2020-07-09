package dm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../dataForms"
	"../utils"
)

func getConvo(sender, receiver string) dataForms.Conversation {
	page := fmt.Sprintf("/message/direct/%v/%v", sender, receiver)
	resp, err := http.Get(BaseUrl + page)
	utils.Check(err)
	defer resp.Body.Close()

	var convo dataForms.Conversation
	respBody, err := ioutil.ReadAll(resp.Body)
	utils.Check(err)

	json.Unmarshal(respBody, &convo)
	return convo
}
