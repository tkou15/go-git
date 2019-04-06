package main

import (
	"fmt"

	"gopkg.in/src-d/go-git.v4"

	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

// Example of how to:
// - Clone a repository into memory
// - Get the HEAD reference
// - Using the HEAD reference, obtain the commit this reference is pointing to
// - Using the commit, obtain its history and print it
func main() {
	// Clones the given repository, creating the remote, the local branches
	// and fetching the objects, everything in memory:
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: "https://github.com/tkou15/go-example.git",
	})
	if err != nil {
		print(err)
	}

	// ... retrieves the branch pointed by HEAD
	ref, err := r.Head()
	if err != nil {
		print(err)
	}

	// ... retrieves the commit history
	cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		print(err)
	}

	// ... just iterates over the commits, printing it
	err = cIter.ForEach(func(c *object.Commit) error {
		fmt.Println(c)

		return nil
	})

	if err != nil {
		print(err)
	}

}
