package usr

import (
	"reflect"
	"testing"
)

func TestGetUserList(t *testing.T) {
	u := GetUserList()
	if !reflect.DeepEqual(Usernames, u) {
		t.Errorf("Expecting %v but instead received %v", Usernames, u)
	}
}
