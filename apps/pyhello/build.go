package main

import (
	"log"
	"path/filepath"

	"github.com/joshuasprow/please-practice/packages/scripts"
	"github.com/magefile/mage/sh"
)

func buildPyhello() error {
	src, err := sh.Output("plz", "build", "//apps/pyhello:pyhello")
	if err != nil {
		return err
	}

	dst := filepath.Join("dist", "pyhello")

	return scripts.CopyFile(src, dst)
}

func copyGreetingsFile() error {
	dst := filepath.Join("dist", "greetings.json")
	src := filepath.Join("..", "..", "data", "greetings.json")

	return scripts.CopyFile(src, dst)
}

// captures output of plz build command (the .pex file's path) and moves it
// to the dist directory
func main() {
	if err := buildPyhello(); err != nil {
		log.Fatal(err)
	}
	if err := copyGreetingsFile(); err != nil {
		log.Fatal(err)
	}
}
