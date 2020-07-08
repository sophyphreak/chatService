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
//	- add to the Channels slice
// - if not unique return error

// AddMessage function
//	- add sent message to Messages slice

// GetMessages function
//	- takes in channel name
//	- returns slice of all messages in that channel

// clearChannels function
//	- sets Channels to empty slice

// GetChannelNames function
//	- returns slice of all channel names
