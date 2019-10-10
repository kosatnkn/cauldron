package project

import (
	"fmt"
	"os"

	"github.com/kosatnkn/cauldron/errors"
	"github.com/kosatnkn/cauldron/repository"
)

// Create creates a new project using `Catalyst` as the base.
// This will panic if it encounters an error.
func Create(name string, modulePart string, tag string) {

	// get project dir
	dir, err := getProjectDir(name)
	if err != nil {
		errors.Handle(err)
	}

	// get module
	module := getModule(modulePart, name)

	// TODO: remove later
	fmt.Println("Module:", module)

	// // create base
	// err = createBase(dir, tag)
	// if err != nil {
	// 	errors.Handle(err)
	// }

	// // clean
	// err = clean(dir)
	// if err != nil {
	// 	errors.Handle(err)
	// }

	// configure
	err = configure(name, module, dir)
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

	// TODO: need to remove later
	path := "/home/kosala/Development/temp"
	fmt.Println("Using temp path", path)

	// // get current location
	// path, err := os.Getwd()
	// if err != nil {
	// 	return "", err
	// }

	// define directory to clone in to
	dir := path + string(os.PathSeparator) + name

	return dir, nil
}

// getModule returns the module name using module part and project name.
func getModule(modulePart string, name string) string {

	// remove all tailing slashes from module part
	for {

		if modulePart[len(modulePart)-1:] != "/" {
			break
		}

		modulePart = modulePart[0 : len(modulePart)-1]
	}

	return fmt.Sprintf("%s/%s", modulePart, name)
}
