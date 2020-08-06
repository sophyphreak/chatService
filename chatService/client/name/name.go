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

	"../dm"
)

//Username holds what a username consist of
type Username struct {
	Username string `json:"username"`
}

//GetUsername gets username from the client
func GetUsername() string {
	var currentUsername string
Welcome:
	fmt.Println("Thank you for using our service!")
	fmt.Print("Are you a returning user? (yes or no) ")
	in := bufio.NewReader(os.Stdin)
	answer, _ := in.ReadString('\n')
	answer = strings.TrimSpace(answer)
	answer = strings.ToLower(answer)

	if answer == "yes" {
		fmt.Print("Please enter your username: ")
		currentUsername, _ := in.ReadString('\n')
		currentUsername = strings.TrimSpace(currentUsername)
		userSlice := dm.GetUserNames()
		if len(userSlice) == 0 {
			fmt.Println("There have been no usernames created.")
			answer = "no"
		}
		for i := 0; i < len(userSlice); i++ {
			if userSlice[i] == currentUsername {
				fmt.Println("We have found your username! Welcome back", currentUsername+"!")
				return currentUsername
			} else {
				fmt.Print("We could not find your username. Would you like to try again? ")
				response, _ := in.ReadString('\n')
				response = strings.TrimSpace(response)
				response = strings.ToLower(response)
				if response == "yes" {
					goto Welcome
				} else {
					answer = "no"
				}
			}
		}
	}
	if answer == "no" {
	Start:
		fmt.Print("Please create your username: ")

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

		currentUsername = secondUsername.Username

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
	}
	if answer != "yes" && answer != "no" {
		fmt.Println("You did not enter a valid answer. Try again.")
		goto Welcome
	}
	return currentUsername
}
