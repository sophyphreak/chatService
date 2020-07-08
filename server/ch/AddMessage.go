package ch

import (
	"fmt"

	"../msg"
)

// AddMessage adds Message to Messages to channel struct
func AddMessage(newMessage msg.Message) {

	var ch Channel
	ch.Messages = append(ch.Messages, newMessage)

	fmt.Println(ch.Messages)

}
