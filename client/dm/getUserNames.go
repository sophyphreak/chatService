package dm

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"../utils"
)

//retrieve list of usernames from server
func GetUserNames() []string {
	//get the response from server
	resp, err := http.Get(utils.URL + "/username")
	utils.Check(err)
	defer resp.Body.Close()

	users := make([]string, 0)
	//read response and convert to byte
	respBody, err := ioutil.ReadAll(resp.Body)
	utils.Check(err)

	//convert to list
	json.Unmarshal(respBody, &users)

	return users
}
