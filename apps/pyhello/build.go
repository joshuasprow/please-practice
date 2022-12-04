package main

import (
	"io"
	"log"
	"os"

	"github.com/magefile/mage/sh"
)

// captures output of plz build command (the .pex file's path) and moves it
// to the dist directory
func main() {
	plzpath, err := sh.Output("plz", "build", "//apps/pyhello:pyhello")
	if err != nil {
		log.Fatal(err)
	}

	os.Mkdir("dist", 0755)

	plzfile, err := os.Open(plzpath)
	if err != nil {
		log.Fatal(err)
	}
	defer plzfile.Close()

	dstfile, err := os.Create("dist/pyhello")
	if err != nil {
		log.Fatal(err)
	}
	defer dstfile.Close()

	if err := dstfile.Chmod(0755); err != nil {
		log.Fatal(err)
	}

	if _, err := io.Copy(dstfile, plzfile); err != nil {
		log.Fatal(err)
	}
}
