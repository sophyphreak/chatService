package usr

import "../cnv"

// Users is a slice of all Users
var Users []User

// User is a user
type User struct {
	Username      string              `json:"username"`
	Conversations []*cnv.Conversation `json:"conversations"`
}
