package project

import (
	"fmt"
	"os"

	"github.com/kosatnkn/cauldron/errors"
	"github.com/kosatnkn/cauldron/log"
	"github.com/kosatnkn/cauldron/repository"
)

// Create creates a new project using `Catalyst` as the base.
// This will panic if it encounters an error.
func Create(name string, modulePrefix string, tag string) {

	// get project dir
	dir, err := getProjectDir(name)
	if err != nil {
		errors.Handle(err)
	}

	// get module
	// module := getModule(modulePrefix, name)

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

	// index
	idx, err := index(dir)
	if err != nil {
		errors.Handle(err)
	}

	// configure
	err = configure(name, modulePrefix, idx)
	if err != nil {
		errors.Handle(err)
	}
}

// createBase creates the project base by cloning `Catalyst`
func createBase(dir string, tag string) error {

	m := fmt.Sprintf("\nProject is created in %s", dir)
	log.Info(m)

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

	// TODO: need to remove later
	path := "/home/kosala/Development/temp"
	log.Warn("\nUsing temp path " + path)

	// // get current location
	// path, err := os.Getwd()
	// if err != nil {
	// 	return "", err
	// }

	// define directory to clone in to
	dir := path + string(os.PathSeparator) + name

	return dir, nil
}

// getModule returns the module name using module prefix and project name.
func getModule(modulePrefix string, name string) string {

	// remove all tailing slashes from module prefix
	for {

		if modulePrefix[len(modulePrefix)-1:] != "/" {
			break
		}

		modulePrefix = modulePrefix[0 : len(modulePrefix)-1]
	}

	return fmt.Sprintf("%s/%s", modulePrefix, name)
}
