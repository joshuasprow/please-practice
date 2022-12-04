package greetings

import "testing"

func TestGreeting(t *testing.T) {
	greeting, err := Greeting()
	if err != nil {
		t.Fatal(err)
	}

	if greeting == "" {
		t.Fatal("Greeting was empty")
	}
}
