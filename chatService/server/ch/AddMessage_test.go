package ch

import (
	"reflect"
	"testing"

	"../msg"
)

// Tests to see if adds Message correctly
func TestAddMessage(t *testing.T) {

	channelName := "The Bois"

	newMessage := msg.Message{
		Username: "jazzyjazz0713",
		Body:     "I am sick and tired of being sick and tired.",
	}
	messageSlice := []msg.Message{newMessage}

	got := AddMessage(newMessage, channelName)

	expected := Channel{channelName, messageSlice}

	if got.Name != expected.Name {
		t.Errorf("Got %v Expected %v", got.Name, expected.Name)
	}
	if !reflect.DeepEqual(got.Messages, expected.Messages) {
		t.Errorf("Got %v Expected %v", got.Messages, expected.Messages)
	}

	clearChannels()
}
