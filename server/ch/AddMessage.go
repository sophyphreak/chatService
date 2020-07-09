package ch

import (
	"../msg"
)

// AddMessage adds Message to Messages to channel struct
func AddMessage(newMessage msg.Message, name string) Channel {

	for indx, chnl := range Channels {
		if chnl.Name == name {
			Channels[indx].Messages = append(Channels[indx].Messages, newMessage)
		}
	}

	messages := make([]msg.Message, 0)
	messages = append(messages, newMessage)

	newChnl := Channel{name, messages}
	Channels = append(Channels, &newChnl)

	return newChnl
}
