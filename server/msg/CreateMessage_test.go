package msg

import "testing"

func TestCreateMessage(t *testing.T) {

	t.Run("some input", func(t *testing.T) {
		userName := "bhehar"
		body := "Life is, but a shit show"

		got := CreateMessage(userName, body)
		expected := Message{userName, body}

		if got != expected {
			t.Errorf("got %v expected %v", got, expected)
		}
	})
	t.Run("some input", func(t *testing.T) {
		userName := "bhehar"
		body := " "

		got := CreateMessage(userName, body)
		expected := Message{userName, body}

		if got != expected {
			t.Errorf("got %v expected %v", got, expected)
		}
	})

}
