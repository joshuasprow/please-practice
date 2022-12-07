package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5/plumbing/format/gitignore"
)

func newPatternMatcher(ignorePatterns ...string) gitignore.Matcher {
	patterns := []gitignore.Pattern{}

	for _, p := range ignorePatterns {
		patterns = append(patterns, gitignore.ParsePattern(p, nil))
	}

	return gitignore.NewMatcher(patterns)
}

// walks filepath and finds filenames ending with .ipynb. pass glob patterns to
// ignore files/directories
func findIpynbPaths(
	root string,
	ignorePatterns ...string,
) (
	[]string,
	error,
) {
	matcher := newPatternMatcher(ignorePatterns...)
	paths := []string{}

	if err := filepath.WalkDir(
		root,
		func(path string, d fs.DirEntry, err error) error {
			if matcher.Match([]string{path}, d.IsDir()) {
				return filepath.SkipDir
			}
			if filepath.Ext(path) == ".ipynb" {
				paths = append(paths, path)
			}
			return nil
		},
	); err != nil {
		return nil, err
	}

	return paths, nil
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

func readIpynbFile(path string) (ipynbData, error) {
	file, err := os.Open(path)
	if err != nil {
		return ipynbData{}, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return ipynbData{}, err
	}

	ipynb := ipynbData{}

	if err := json.Unmarshal(data, &ipynb); err != nil {
		return ipynbData{}, err
	}

	return ipynb, nil
}

func scrubOutputData(ipynb ipynbData) ipynbData {
	for i := range ipynb.Cells {
		ipynb.Cells[i].ExecutionCount = 0
		ipynb.Cells[i].Outputs = nil
	}

	return ipynb
}

func writeIpynbFile(path string, ipynb ipynbData) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.MarshalIndent(ipynb, "", "  ")
	if err != nil {
		return err
	}

	if _, err := file.Write(data); err != nil {
		return err
	}

	return nil
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

		scrubbed := scrubOutputData(ipynb)
		scrubbedPath := fmt.Sprintf("%s.scrubbed", path)

		if err := writeIpynbFile(scrubbedPath, scrubbed); err != nil {
			panic(err)
		}
	}

	os.Exit(1)
}
