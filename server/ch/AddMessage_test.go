package ch

import (
	"testing"

	"../msg"
)

// Tests to see if adds Message correctly
func TestAddMessage(t *testing.T) {

	beforeAddMessage := len(Channels)

	channelName := "The Bois"

	newMessage := msg.Message{
		"jazzyjazz0713",
		"I am sick and tired of being sick and tired.",
	}
	AddMessage(newMessage, channelName)

	afterAddMessage := len(Channels)

	if beforeAddMessage == afterAddMessage {
		t.Errorf("Chanel Message should have updated")
	}

	clearChannels()
}
