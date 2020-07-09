package dataForms

type Conversation struct {
	User1    string    `json:"user1"`
	User2    string    `json:"user2"`
	Messages []Message `json:"messages"`
}
