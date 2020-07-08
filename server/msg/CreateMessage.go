package msg

func CreateMessage(userName, body string) Message {
	return Message{userName, body}
}
