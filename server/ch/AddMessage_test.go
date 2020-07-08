package ch

import (
	"fmt"
	"testing"

	"../msg"
)

// Tests to see if adds Message correctly
func TestAddMessage(t *testing.T) {

	var ch Channel

	fmt.Println(ch.Messages)
	empty := len(ch.Messages)

	newMessage := msg.Message{
		"jazzyjazz0713",
		"I am sick and tired of being sick and tired.",
	}
	AddMessage(newMessage)

	hopeful := len(ch.Messages)
	fmt.Println(ch.Messages)

	if empty == hopeful {
		t.Errorf("Chanel Message should have updated")
	}

}
