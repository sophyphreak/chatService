package ch

import "../msg"

// AddChannel adds a new channel or returns channel with same name
func AddChannel(name string) Channel {
	for _, c := range Channels {
		ch := *c
		if ch.Name == name {
			return ch
		}
	}
	ch := Channel{name, make([]msg.Message, 0)}
	Channels = append(Channels, &ch)
	return ch
}
