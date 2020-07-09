package dm

import (
	"fmt"

	"../utils"
)

const (
	BaseUrl = "http://localhost:10000"
)

//prints error when user enter invalid message
func printMessageError() {
	fmt.Println()
	fmt.Println("Please enter a message!")
	fmt.Println()
}

func RunDM(sender string) {
	availableUsers := getUserNames()
	utils.PrintStringSlice(availableUsers)
ChooseReceiver:
	receiver := utils.GetInput("Send message to: ")
	if !utils.StringInSlice(receiver, availableUsers) {
		fmt.Println("That user is not available")
		goto ChooseReceiver
	}
	quit := make(chan byte)
	go listenForDm(sender, receiver, quit)
GetMessage:
	message := utils.GetInput("")
	if message == "quit" {
		quit <- 'q'
		return
	}
	if message == "" {
		printMessageError()
		goto GetMessage
	}
	sendMessage(sender, receiver, message)
	goto GetMessage

}
