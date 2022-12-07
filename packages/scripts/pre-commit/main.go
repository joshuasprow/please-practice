package main

func main() {
	ignorePatterns, err := readGitignoreFile(".gitignore")
	if err != nil {
		panic(err)
	}

	if err := scrubIpynbFiles(ignorePatterns...); err != nil {
		panic(err)
	}
}
