package dm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../utils"
)

func getConvo(sender, receiver string) conversation {
	page := fmt.Sprintf("/message/direct/%v/%v", sender, receiver)
	resp, err := http.Get(utils.URL + page)
	utils.Check(err)
	defer resp.Body.Close()

	var convo conversation
	respBody, err := ioutil.ReadAll(resp.Body)
	utils.Check(err)

	json.Unmarshal(respBody, &convo)
	return convo
}
