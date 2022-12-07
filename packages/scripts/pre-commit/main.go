package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/go-git/go-git/v5/plumbing/format/gitignore"
	"github.com/magefile/mage/sh"
)

func newPatternMatcher(ignorePatterns ...string) gitignore.Matcher {
	patterns := []gitignore.Pattern{}

	for _, p := range ignorePatterns {
		patterns = append(patterns, gitignore.ParsePattern(p, nil))
	}

	return gitignore.NewMatcher(patterns)
}

func readGitignoreFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || line[0] == '#' {
			continue
		}

		lines = append(lines, line)
	}

	return lines, nil
}

func gitAddPaths(paths []string) error {
	args := []string{"add"}
	args = append(args, paths...)

	return sh.RunV("git", args...)
}

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
