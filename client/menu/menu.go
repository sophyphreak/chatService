package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"../addch"

	"../dm"
	"../joinch"
)

//Menu displays the users choices within the chat service
func Menu(username string) {

	var response string
	in := bufio.NewReader(os.Stdin)

	for {

		fmt.Println()
		fmt.Println("\t\tWelcome to our Chat Service", username+"!\n")
		fmt.Println("Select one of the choices below!")
		fmt.Println("1. Create a new channel")
		fmt.Println("2. Join an existing channel")
		fmt.Println("3. Start a direct message with another user")
		fmt.Println("4. Quit")

		fmt.Print("Select a number 1-4: ")
		response, _ = in.ReadString('\n')
		response = strings.TrimSpace(response)

		switch response {

		case "1":
			addch.Begin()
		case "2":
			joinch.Begin(username)
		case "3":
			dm.RunDM(username)
		case "4":
			fmt.Println("Thank you", username, "for chatting with us!")
			os.Exit(1)
		default:
			fmt.Println("Invalid response. Please try again.")
			continue
		}
	}
}
