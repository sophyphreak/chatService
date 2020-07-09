package name

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

//Username holds what a username consist of
type Username struct {
	Username string `json:"username"`
}

//GetUsername gets username from the client
func GetUsername() string {

	fmt.Println("Thank you for using our service!")
Start:
	fmt.Print("Please create your username: ")

	in := bufio.NewReader(os.Stdin)
	name, _ := in.ReadString('\n')
	name = strings.TrimSpace(name)
	matched, _ := regexp.MatchString(`\w`, name)

	if matched == false {
		fmt.Println("Your username did not contain any characters. Try again")
		fmt.Println()
		goto Start
	}
	username := Username{name}
	jsonValue, _ := json.Marshal(username)
	resp, _ := http.Post("http://localhost:10000/username", "application/json", bytes.NewBuffer(jsonValue))

	var secondUsername Username
	respBody, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(respBody, &secondUsername)

	currentUsername := secondUsername.Username

	if currentUsername != name {
		fmt.Println("The username", name, "has been taken. Your suggested username is", currentUsername)
		fmt.Print("Would you like to use the suggested username? ")
		response, _ := in.ReadString('\n')
		response = strings.ToLower(response)
		response = strings.TrimSpace(response)
		if response == "no" {
			goto Start
		} else {
			fmt.Println("Your username has successfully been created")
		}
	} else {
		fmt.Println("Your username has successfully been created")
	}

	return currentUsername

}
