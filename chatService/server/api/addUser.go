package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"../usr"
)

type username struct {
	Username string `json:"username"`
}

func addUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var u username
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.Unmarshal(reqBody, &u)

	username := usr.AddUser(u.Username)

	res := make(map[string]string)
	res["username"] = username
	json.NewEncoder(w).Encode(res)
}
