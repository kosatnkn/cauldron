package project

import (
	"os"

	"github.com/kosatnkn/cauldron/errors"
	"github.com/kosatnkn/cauldron/repository"
)

// Create creates a new project using `Catalyst` as the base.
// This will panic if it encounters an error.
func Create(name string, tag string) {

	// get project dir
	dir, err := getProjectDir(name)
	if err != nil {
		errors.Handle(err)
	}

	// create base
	err = createBase(dir, tag)
	if err != nil {
		errors.Handle(err)
	}

	// clean
	err = clean(dir)
	if err != nil {
		errors.Handle(err)
	}

}

// createBase creates the project base by cloning `Catalyst`
func createBase(dir string, tag string) error {

	// clone
	r, err := repository.Clone(dir)
	if err != nil {
		return err
	}

	// checkout tag
	err = repository.Checkout(r, tag)
	if err != nil {
		return err
	}

	return nil
}

// getProjectDir gives the directory path to create the repository.
func getProjectDir(name string) (string, error) {

	// get current location
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// define directory to clone in to
	dir := path + string(os.PathSeparator) + name

	return dir, nil
}
