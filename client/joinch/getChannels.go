package joinch

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"../utils"
)

func getChannels() []string {
	var channels []string
	res, err := http.Get(utils.URL + "/channel")
	if err != nil {
		log.Fatal(err)
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(resBody, &channels)
	return channels
}
