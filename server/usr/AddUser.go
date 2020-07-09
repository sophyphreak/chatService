package usr

import (
	"strconv"
	"time"
)

// AddUser sees if a user is unique and then updates Usernames array
func AddUser(user string) string {

	for i := 0; i < len(Usernames); i++ {
		if Usernames[i] == user {
			t := time.Now().Unix()
			unixTimeStamp := strconv.FormatInt(t, 10)
			user = user + unixTimeStamp
			break
		}
	}

	Usernames = append(Usernames, user)

	return user
}
