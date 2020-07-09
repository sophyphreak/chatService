package dm

import (
	"bytes"
	"encoding/json"
	"net/http"

	"../dataForms"
	"../utils"
)

func sendMessage(sender, receiver, body string) {

	dm := dataForms.DirectMessage{Sender: sender, Receiver: receiver, Body: body}
	page := "/message/direct"
	jsonValue, err := json.Marshal(dm)
	utils.Check(err)
	resp, err := http.Post(BaseUrl+page, "application/json", bytes.NewBuffer(jsonValue))
	utils.Check(err)
	resp.Body.Close()
}
