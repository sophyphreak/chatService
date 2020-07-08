package msg

// Message is a single message
type Message struct {
	Username string `json:"username"`
	Body     string `json:"body"`
}
