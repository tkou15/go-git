package main

import (
	"os"

	git "gopkg.in/src-d/go-git.v4"
)

func main() {
	git.PlainClone("/tmp/foo", false, &git.CloneOptions{
		URL:      "https://github.com/src-d/go-git",
		Progress: os.Stdout,
	})

}
