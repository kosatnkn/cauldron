package project

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/kosatnkn/cauldron/config"
	"github.com/kosatnkn/cauldron/errors"
	"github.com/kosatnkn/cauldron/log"
	"github.com/kosatnkn/cauldron/repository"
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

// createBase creates the project base by cloning the base repository.
func createBase(dir string, cfg *config.Config) error {

	m := fmt.Sprintf("Project is created in %s", dir)
	log.Note(m)

	// clone
	r, err := repository.Clone(dir, cfg.Base.Repo)
	if err != nil {
		return err
	}

	// checkout tag
	tag := cfg.Base.Version
	if tag == "" {
		tag = cfg.Base.MaxVersion
	}
	if isIllegalVersion(cfg, tag) {
		return fmt.Errorf(`Creating a new project using '%s' of '%s' is not supported by Cauldron '%s'`, tag, cfg.Base.Repo, cfg.Cauldron.Version)
	}

	err = repository.Checkout(r, tag)
	if err != nil {
		return err
	}

	return nil
}

// isIllegalVersion checks whether the provided version is supported by Cauldron.
func isIllegalVersion(cfg *config.Config, version string) bool {

	v, _ := strconv.Atoi(strings.Join(strings.Split(version[1:], "."), ""))
	min, _ := strconv.Atoi(strings.Join(strings.Split(cfg.Base.MinVersion[1:], "."), ""))
	max, _ := strconv.Atoi(strings.Join(strings.Split(cfg.Base.MaxVersion[1:], "."), ""))

	return v >= min && v <= max
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
//
// Will only allow lower case letters, numbers and `-` character.
// Upper case characters will be converted to lower case characters and
// any other non permitted characters will be removed.
func simplifyName(name string) string {

	reg, _ := regexp.Compile("[^a-zA-Z0-9-]+")

	name = reg.ReplaceAllString(name, "")

	return strings.ToLower(name)
}
