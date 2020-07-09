package ch

import (
	"testing"

	"../msg"
)

func TestClearChannels(t *testing.T) {
	msg1 := msg.Message{Username: "james", Body: "What's up dude"}
	msgSlice := make([]msg.Message, 1)
	msgSlice[0] = msg1

	ch1 := Channel{"testChannel", msgSlice}
	Channels = append(Channels, &ch1)
	clearChannels()

	got := len(Channels)
	expected := 0

	if got != expected {
		t.Errorf("Expected %v got %v", expected, got)
	}
}
