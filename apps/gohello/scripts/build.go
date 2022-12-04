package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

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
	dst, err := filepath.Abs(filepath.Join("dist", "greetings.json"))
	if err != nil {
		return err
	}

	src, err := filepath.Abs(
		filepath.Join("..", "..", "data", "greetings.json"),
	)
	if err != nil {
		return err
	}

	fmt.Printf("copying %s to %s\n", src, dst)

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := buildGoHello(); err != nil {
		log.Fatal(err)
	}
	if err := copyGreetingsFile(); err != nil {
		log.Fatal(err)
	}
}
