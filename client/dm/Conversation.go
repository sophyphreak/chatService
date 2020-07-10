package dm

import (
	"../msg"
)

type conversation struct {
	User1    string        `json:"user1"`
	User2    string        `json:"user2"`
	Messages []msg.Message `json:"messages"`
}
