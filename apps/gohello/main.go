package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joshuasprow/please-practice/packages/greetings"
)

func main() {
	path := os.Getenv("GREETINGS_PATH")

	greeting, err := greetings.Greeting(path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s World\n", greeting)
}
