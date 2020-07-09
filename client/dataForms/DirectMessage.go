package dataForms

type DirectMessage struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Body     string `json:"body"`
}
