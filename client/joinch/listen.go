package joinch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func listen(channelName string, quit chan struct{}) {
	var messages []string
	select {
	case <-quit:
		return
	default:
		res, err := http.Get("http://localhost:10000/message/channel/" + channelName)
		if err != nil {
			log.Fatal(err)
		}
		resBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		var resMessages []string
		json.Unmarshal(resBody, &resMessages)
		if len(resMessages) < len(messages) {
			for i := len(messages); i < len(resMessages); i++ {
				fmt.Println(resMessages[i])
				messages = resMessages
			}
		}
		time.Sleep(time.Duration(250) * time.Millisecond)
	}
}
