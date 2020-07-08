package msg

// Message is a single message
type Message struct {
	Username string `json:"username"`
	Body     string `json:"body"`
}

// CreateMessage function
//	- takes in username and body
//	- returns new message
