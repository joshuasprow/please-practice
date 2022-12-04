package greetings

import (
	"path/filepath"
	"testing"
)

func TestGreeting(t *testing.T) {
	greeting, err := Greeting(
		filepath.Join("..", "..", "data", "greetings.json"),
	)
	if err != nil {
		t.Fatal(err)
	}

	if greeting == "" {
		t.Fatal("Greeting was empty")
	}
}
