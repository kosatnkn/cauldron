package project

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/kosatnkn/cauldron/v2/config"
	"github.com/kosatnkn/cauldron/v2/errors"
	"github.com/kosatnkn/cauldron/v2/log"
	"github.com/kosatnkn/cauldron/v2/repository"
)

// Create creates a new project using the base project.
func Create(cfg *config.Config) {

	simpleName := simplifyName(cfg.Project.Name)

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
	err = clean(dir, cfg.Base.RemoveDirs, cfg.Base.RemoveFiles)
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

// createBase creates the project base by cloning the base repository.
func createBase(dir string, cfg *config.Config) error {

	m := fmt.Sprintf("Project is created in %s", dir)
	log.Note(m)

	// clone
	r, err := repository.Clone(dir, cfg.Base.Repo)
	if err != nil {
		return err
	}

	err = repository.Checkout(r, cfg.Base.Version, cfg.Base.MaxVersion)
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
func getModule(prefix, name string) string {

	// remove all tailing slashes from module prefix
	for {

		if prefix[len(prefix)-1:] != "/" {
			break
		}

		prefix = prefix[0 : len(prefix)-1]
	}

	return fmt.Sprintf("%s/%s", prefix, name)
}

// simplifyName removes all illegal characters and lowercase the name.
//
// Will only allow lower case letters, numbers and `-` character.
// Upper case characters will be converted to lower case characters and
// any other non permitted characters will be removed.
func simplifyName(name string) string {

	reg, _ := regexp.Compile("[^a-zA-Z0-9-]+")

	name = reg.ReplaceAllString(name, "")

	return strings.ToLower(name)
}
