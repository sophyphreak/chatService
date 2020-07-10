package joinch

import (
	"fmt"
)

// Begin is the entry point to join channel functionality
func Begin(username string) {
	channels := getChannels()
	currentChannel := chooseChannel(channels)
	quit := make(chan struct{})
	go listen(currentChannel, quit)
	fmt.Println("If you would like to leave this channel type \"quit\"")
	fmt.Println()
	inputMessages(currentChannel, username)
	quit <- struct{}{}
}
