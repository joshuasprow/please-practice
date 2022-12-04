package main

import (
	"fmt"
	"log"

	"github.com/joshuasprow/please-practice/packages/greetings"
)

func main() {
	greeting, err := greetings.Greeting()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s World\n", greeting)
}
