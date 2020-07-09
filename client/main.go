package main

import (
	"fmt"

	"./dm"
)

func main() {
	sender := "ballzie"
	receiver := "jazzySmith"

	convo := dm.SendMessage(sender, receiver, "I'm having a dope ass day!")
	fmt.Println(convo)
}
