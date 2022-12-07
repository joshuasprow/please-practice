package main

import (
	"fmt"
)

func main() {
	ignorePatterns, err := readGitignoreFile(".gitignore")
	if err != nil {
		panic(err)
	}

	paths, err := findIpynbPaths(".", ignorePatterns...)
	if err != nil {
		panic(err)
	}

	for _, path := range paths {
		ipynb, err := readIpynbFile(path)
		if err != nil {
			panic(err)
		}

		if !ipynb.HasOutputs() {
			continue
		}

		scrubbed := scrubOutputData(ipynb)

		if err := writeIpynbFile(path, scrubbed); err != nil {
			panic(err)
		}

		fmt.Printf("scrubbed %s\n", path)
	}

	if err := gitAddPaths(paths); err != nil {
		panic(err)
	}
}
