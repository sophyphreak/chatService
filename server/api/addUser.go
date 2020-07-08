package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type username struct {
	Username string `json:"username"`
}

func addUser(w http.ResponseWriter, r *http.Request) {
	var u username
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.Unmarshal(reqBody, &u)
	username := u // will be replaced
	json.NewEncoder(w).Encode(username)
}
