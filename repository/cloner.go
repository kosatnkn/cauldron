package repository

import (
	"fmt"
	"os"

	"gopkg.in/src-d/go-git.v4"
)

const catalystURL string = "https://github.com/kosatnkn/catalyst.git"

// Clone clones `Catalyst` to the given directory.
func Clone(dir string) (*git.Repository, error) {

	fmt.Println("Cloning Catalyst from", catalystURL)

	opts := &git.CloneOptions{
		URL:      catalystURL,
		Progress: os.Stdout,
	}

	r, err := git.PlainClone(dir, false, opts)
	if err != nil {
		return nil, err
	}

	return r, nil
}
