package cnv

import "../msg"

// Conversation is a private conversation between two Users
type Conversation struct {
	User1    string        `json:"user1"`
	User2    string        `json:"user2"`
	Messages []msg.Message `json:"messages"`
}
