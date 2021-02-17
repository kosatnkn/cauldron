package config

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	e "github.com/kosatnkn/cauldron/errors"
)

// Validate validates the populated config struct.
func Validate(cfg *Config) {

	var err error

	err = isNameValid(cfg.Project.Name)
	if err != nil {
		e.Handle(err)
	}

	err = isNameSpaceValid(cfg.Project.Namespace)
	if err != nil {
		e.Handle(err)
	}

	// checkout tag
	if !isVersionInRange(cfg.Base.Version, cfg.Base.MinVersion, cfg.Base.MaxVersion) {
		err = fmt.Errorf(`Cauldron(%s) supports '%s' to '%s' of '%s', cannot create project using '%s'`,
			cfg.Cauldron.Version,
			cfg.Base.MinVersion,
			cfg.Base.MaxVersion,
			cfg.Base.Repo,
			cfg.Base.Version)

		e.Handle(err)
	}

}

// isNameValid checks whether name is valid.
func isNameValid(name string) error {

	if name == "" {
		return errors.New("Project name cannot be empty")
	}

	return nil
}

// isNameSpaceValid checks whether namespace is valid.
func isNameSpaceValid(namespace string) error {

	pattern := regexp.MustCompile(`^[a-zA-Z0-9_\-\/.]*$`).MatchString

	if namespace == "" {
		return errors.New("NameSpace cannot be empty")
	}

	if !pattern(namespace) {
		return errors.New("Namespace contain illegal characters")
	}

	return nil
}

// isVersionInRange checks whether the provided version is in the supported version range.
func isVersionInRange(version, min, max string) bool {

	if version == "" {
		return true
	}

	v, _ := strconv.Atoi(strings.Join(strings.Split(version[1:], "."), ""))
	mn, _ := strconv.Atoi(strings.Join(strings.Split(min[1:], "."), ""))
	mx, _ := strconv.Atoi(strings.Join(strings.Split(max[1:], "."), ""))

	return v >= mn && v <= mx
}
