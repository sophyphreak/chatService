package ch

import (
	"testing"
)

func TestAddChannel(t *testing.T) {
	defer clearChannels()
	clearChannels()
	name := "bob channel"
	AddChannel(name)
	ch := Channels[0]
	if name != ch.Name {
		t.Errorf("Expected channel name to be %s but instead received %s", name, ch.Name)
	}
	if len(ch.Messages) != 0 {
		t.Errorf("Expected channel messages to be length zero but instead received length %d", len(ch.Messages))
	}
}
