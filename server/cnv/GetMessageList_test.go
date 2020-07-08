package cnv

import (
	"testing"

	"../msg"
)

func TestGetMessages(t *testing.T) {
	defer func() { Conversations = Conversations[:0] }()

	var m msg.Message
	m.Username = "bali"
	m.Body = "wzzzp"

	var c Conversation
	c.User1 = "bali"
	c.User2 = "jasmine"
	c.Messages = append(c.Messages, m)

	Conversations = append(Conversations, c)

}
