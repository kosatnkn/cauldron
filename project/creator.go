package project

import (
	"github.com/kosatnkn/cauldron/repository"
)

// Create creates a new project using `Catalyst` as the base.
// This will panic if it encounters an error.
func Create(name string, tag string) {

	// create base
	err := createBase(name, tag)
	if err != nil {
		panic(err)
	}
}

// createBase creates the project base by cloning `Catalyst`
func createBase(name string, tag string) error {

	// clone
	r, err := repository.Clone(name)
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
