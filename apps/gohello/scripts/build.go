package main

import (
	"log"
	"path/filepath"

	"github.com/joshuasprow/please-practice/packages/scripts"
	"github.com/magefile/mage/sh"
)

func buildGoHello() error {
	log.Println("building gohello")

	return sh.RunWith(
		map[string]string{"GOOS": "linux", "GOARCH": "amd64"},
		"go", "build",
		"-o", "dist/gohello",
		"main.go",
	)
}

func copyGreetingsFile() error {
	dst := filepath.Join("dist", "greetings.json")
	src := filepath.Join("..", "..", "data", "greetings.json")

	return scripts.CopyFile(src, dst)
}

func main() {
	if err := buildGoHello(); err != nil {
		log.Fatal(err)
	}
	if err := copyGreetingsFile(); err != nil {
		log.Fatal(err)
	}
}
