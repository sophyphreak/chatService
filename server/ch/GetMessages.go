package ch

import "../msg"

// GetMessages returns a slice of all messages in a channel
func GetMessages(channelName string) []msg.Message {
	for _, c := range Channels {
		if c.Name == channelName {
			return c.Messages
		}
	}
	return make([]msg.Message, 0)
}
