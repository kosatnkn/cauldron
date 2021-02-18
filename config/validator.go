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

	err = checkNameValid(cfg.Project.Name)
	if err != nil {
		e.Handle(err)
	}

	err = checkNameSpaceValid(cfg.Project.Namespace)
	if err != nil {
		e.Handle(err)
	}

	err = checkValidSymanticTag(cfg.Base.Version)
	if err != nil {
		e.Handle(err)
	}

	err = checkVersionInRange(cfg)
	if err != nil {
		e.Handle(err)
	}
}

// checkNameValid checks whether name is valid.
func checkNameValid(name string) error {

	if name == "" {
		return errors.New("Project name cannot be empty")
	}

	return nil
}

// checkNameSpaceValid checks whether namespace is valid.
func checkNameSpaceValid(namespace string) error {

	pattern := regexp.MustCompile(`^[a-zA-Z0-9_\-\/.]*$`).MatchString

	if namespace == "" {
		return errors.New("NameSpace cannot be empty")
	}

	if !pattern(namespace) {
		return errors.New("Namespace contain illegal characters")
	}

	return nil
}

// checkValidSymanticTag checks whether the provided tag is a valid symantic tag
func checkValidSymanticTag(tag string) error {

	if tag == "" {
		return nil
	}

	exp := `^v(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)`

	if !regexp.MustCompile(exp).MatchString(tag) {
		return fmt.Errorf(`'%s' is not a valid symantic tag`, tag)
	}

	return nil
}

// checkVersionInRange checks whether the provided version is in the supported version range.
func checkVersionInRange(cfg *Config) error {

	if cfg.Base.Version == "" {
		return nil
	}

	v, _ := strconv.Atoi(strings.Join(strings.Split(cfg.Base.Version[1:], "."), ""))
	min, _ := strconv.Atoi(strings.Join(strings.Split(cfg.Base.MinVersion[1:], "."), ""))
	max, _ := strconv.Atoi(strings.Join(strings.Split(cfg.Base.MaxVersion[1:], "."), ""))

	if v >= min && v <= max {
		return nil
	}

	return fmt.Errorf(`Cauldron(%s) supports '%s' to '%s' of '%s', cannot create project using '%s'`,
		cfg.Cauldron.Version,
		cfg.Base.MinVersion,
		cfg.Base.MaxVersion,
		cfg.Base.Repo,
		cfg.Base.Version)
}
