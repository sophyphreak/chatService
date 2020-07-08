package ch

import "../msg"

// Channels is a slice of all Channels
var Channels []*Channel

// Channel is a channel
type Channel struct {
	Name     string        `json:"name"`
	Messages []msg.Message `json:"messages"`
}

// CreateChannel function
//	- create channel with name, empty slice
//	- check if channel is unqiue, if not return error
//	- add to the Channels slice

// AddMessage function
//	- add sent message to Messages slice

// GetMessages function
//	- takes in channel name
//	- returns slice of all messages in that channel

// clearChannels function
//	- sets Channels to empty slice

// GetChannelNames function
//	- returns slice of all channel names
