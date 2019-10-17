package project

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/kosatnkn/cauldron/config"
	"github.com/kosatnkn/cauldron/errors"
	"github.com/kosatnkn/cauldron/log"
	"github.com/kosatnkn/cauldron/repository"
)

// Create creates a new project using `Catalyst` as the base.
func Create(cfg *config.Config) {

	simpleName := simplifyName(cfg.Name)

	// get project dir
	dir, err := getProjectDir(simpleName)
	if err != nil {
		errors.Handle(err)
	}

	// create base
	err = createBase(dir, cfg)
	if err != nil {
		errors.Handle(err)
	}

	// clean base
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
	err = configure(cfg, simpleName, idx)
	if err != nil {
		errors.Handle(err)
	}

	// init as new repo
	err = repository.Init(dir)
	if err != nil {
		errors.Handle(err)
	}
}

// createBase creates the project base by cloning `Catalyst`
func createBase(dir string, cfg *config.Config) error {

	m := fmt.Sprintf("Project is created in %s", dir)
	log.Note(m)

	// clone
	r, err := repository.Clone(dir, cfg.Repo)
	if err != nil {
		return err
	}

	// checkout tag
	err = repository.Checkout(r, cfg.Tag)
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

// simplifyName removes all illegal characters and lowercase the name.
func simplifyName(name string) string {

	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")

	name = reg.ReplaceAllString(name, "")

	return strings.ToLower(name)
}
