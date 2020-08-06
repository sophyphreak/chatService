package usr

import (
	"testing"
)

// Tests to see if adds usernames correctly
func TestAddUser1(t *testing.T) {

	username := "Jazzyjazz0713"

	foundUser := AddUser(username)

	if foundUser != username {
		t.Errorf("Got username %s but expected %s", foundUser, username)
	}

}

// Tests to see if it adds timeStamp to username if it already exists
func TestAddUser2(t *testing.T) {

	username := "Jazzyjazz0713"

	foundUser := AddUser(username)

	if foundUser == username {
		t.Errorf("The username did not change")
	}

}
