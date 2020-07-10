package joinch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"../msg"
	"../utils"
)

type response struct {
	Messages []msg.Message `json:"messages"`
}

func listen(channelName string, quit chan struct{}) {
	var messages []msg.Message
Top:
	select {
	case <-quit:
		return
	default:
		res, err := http.Get(utils.URL + "/message/channel/" + channelName)
		if err != nil {
			log.Fatal(err)
		}
		resBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		var r response
		json.Unmarshal(resBody, &r)
		if len(r.Messages) > len(messages) {
			for i := len(messages); i < len(r.Messages); i++ {
				fmt.Println(r.Messages[i].Username+":", r.Messages[i].Body)
				messages = r.Messages
			}
		}
		time.Sleep(time.Duration(250) * time.Millisecond)
		goto Top
	}
}
