package ch

// GetChannelNames returns a slice of Channel names
func GetChannelNames() []string {
	names := make([]string, len(Channels))
	for indx, val := range Channels {
		names[indx] = val.Name
	}
	return names
}
