package ch

import (
	"testing"

	"../msg"
)

func TestGetMessages(t *testing.T) {
	defer clearChannels()
	clearChannels()
	m := msg.Message{Username: "bob", Body: "was here"}
	n := msg.Message{Username: "george", Body: "was also here"}

	name := "bob channel"
	ch := Channel{name, make([]msg.Message, 0)}
	ch.Messages = append(ch.Messages, m)
	ch.Messages = append(ch.Messages, n)

	Channels = append(Channels, &ch)
	messages := GetMessages(name)
	if len(messages) != 2 {
		t.Errorf("Expected to receive messages slice with length of 2 but instead received slice with length of %d", len(messages))
	}
	if m.Username != messages[0].Username {
		t.Errorf("Expected to receive %s as message's username but instead received %s", m.Username, messages[0].Username)
	}
	if m.Body != messages[0].Body {
		t.Errorf("Expected to receive %s as message's username but instead received %s", m.Body, messages[0].Body)
	}
	if n.Username != messages[1].Username {
		t.Errorf("Expected to receive %s as message's username but instead received %s", n.Username, messages[1].Username)
	}
	if n.Body != messages[1].Body {
		t.Errorf("Expected to receive %s as message's username but instead received %s", n.Body, messages[1].Body)
	}
}
