package config

import (
	"errors"
	"regexp"

	e "github.com/kosatnkn/cauldron/errors"
)

// Validate validates the populated config struct.
func Validate(cfg *Config) {

	var err error

	err = isNameValid(cfg.Name)
	if err != nil {
		e.Handle(err)
	}

	err = isNameSpaceValid(cfg.Namespace)
	if err != nil {
		e.Handle(err)
	}
}

// isNameValid checks whether name is valid.
func isNameValid(name string) error {

	if name == "" {
		return errors.New("Name cannot be empty")
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
