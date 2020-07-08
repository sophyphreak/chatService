package ch

import (
	"../msg"
)

// AddMessage adds Message to Messages to channel struct
func AddMessage(newMessage msg.Message, name string) {

	for indx, chnl := range Channels {
		if chnl.Name == name {
			Channels[indx].Messages = append(Channels[indx].Messages, newMessage)
		}
	}

}
