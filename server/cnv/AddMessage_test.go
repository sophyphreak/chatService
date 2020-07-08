package cnv

import (
	"reflect"
	"testing"

	"../msg"
)

func TestAddMessage(t *testing.T) {
	user1 := "bhehar"
	user2 := "andrew"
	newMsg := msg.CreateMessage(user1, "what's up dude")
	messageSlice := make([]msg.Message, 0)
	messageSlice = append(messageSlice, newMsg)

	//adds message to global var Conversations
	AddMessage(user1, user2, newMsg)
	//reset global var Conversations to empty
	defer func() { Conversations = Conversations[:0] }()

	got := Conversations[0]
	expected := Conversation{user1, user2, messageSlice}

	if got.User1 != expected.User1 {
		t.Errorf("Got User 1: %v Expected User 1: %v", got, expected)
	}

	if got.User2 != expected.User2 {
		t.Errorf("Got User 2: %v Expected User 2: %v", got, expected)
	}

	if reflect.DeepEqual(got.Messages, expected.Messages) == false {
		t.Errorf("Got messages: %v Expected message: %v", got, expected)
	}

}
