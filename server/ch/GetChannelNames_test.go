package ch

import (
	"reflect"
	"testing"

	"../msg"
)

func TestGetChannelNames(t *testing.T) {
	defer clearChannels()
	msg1 := msg.Message{Username: "james", Body: "What's up dude"}
	msgSlice := make([]msg.Message, 1)
	msgSlice[0] = msg1

	ch1 := Channel{"testChannel", msgSlice}
	Channels = append(Channels, &ch1)

	got := GetChannelNames()
	expected := []string{"testChannel"}

	if reflect.DeepEqual(got, expected) == false {
		t.Errorf("expected %v got %v", expected, got)
	}

}
