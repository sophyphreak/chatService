package ch

import "../msg"

// Channels is a slice of all Channels
var Channels []Channel

// Channel is a channel
type Channel struct {
	Name     string        `json:"name"`
	Messages []msg.Message `json:"messages"`
}
