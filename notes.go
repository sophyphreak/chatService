package main

// Server

// Users is a slice of all Users
var Users []User

// Channels is a slice of all Channels
var Channels []Channel

// User is a user
type User struct {
	Username      string          `json:"username"`
	Conversations []*Conversation `json:"conversations"`
}

// Conversation is a private conversation between two Users
type Conversation struct {
	User1    string    `json:"user1"`
	User2    string    `json:"user2"`
	Messages []Message `json:"messages"`
}

// Channel is a channel
type Channel struct {
	Name     string    `json:"name"`
	Messages []Message `json:"messages"`
}

// Message is a single message
type Message struct {
	Username string `json:"username"`
	Body     string `json:"body"`
}
