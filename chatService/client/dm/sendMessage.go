package dm

import (
	"bytes"
	"encoding/json"
	"net/http"

	"../utils"
)

func sendMessage(sender, receiver, body string) {

	dm := directMessage{Sender: sender, Receiver: receiver, Body: body}
	fullURL := utils.URL + "/message/direct"
	jsonValue, err := json.Marshal(dm)
	utils.Check(err)
	resp, err := http.Post(fullURL, "application/json", bytes.NewBuffer(jsonValue))
	utils.Check(err)
	resp.Body.Close()
}
