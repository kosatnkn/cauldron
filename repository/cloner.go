package repository

import (
	"os"

	"gopkg.in/src-d/go-git.v4"
)

const catalystURL string = "https://github.com/kosatnkn/catalyst.git"

// Clone clones `Catalyst` to a new directory named after the project name.
func Clone(name string) (*git.Repository, error) {

	// get clone dir
	dir, err := getCloningDir(name)
	if err != nil {
		return nil, err
	}
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

// getCloningDir gives the directory path to clone the repository.
func getCloningDir(name string) (string, error) {

	// get current location
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// define directory to clone in to
	dir := path + string(os.PathSeparator) + name

	return dir, nil
}
