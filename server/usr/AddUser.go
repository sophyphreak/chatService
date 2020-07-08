package usr

import (
	"time"
)

// AddUser sees if a user is unique and then adds to
func AddUser(user string) string {

	for i := 0; i < len(Usernames); i++ {
		if Usernames[i] == user {
			t := time.Now()
			user = user + t.String()
			break
		}
	}

	Usernames = append(Usernames, user)

	return user
}
