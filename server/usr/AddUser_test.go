package usr

import (
	"testing"
)

func TestAddUser1(t *testing.T) {

	username := "Jazzyjazz0713"

	foundUser := AddUser(username)

	if foundUser != username {
		t.Errorf("Got username %s but expected %s", foundUser, username)
	}

}

func TestAddUser2(t *testing.T) {

	username := "Jazzyjazz0713"

	foundUser := AddUser(username)

	if foundUser == username {
		t.Errorf("The username did not change")
	}

}
