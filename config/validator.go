package config

import (
	"errors"

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

	if namespace == "" {
		return errors.New("NameSpace cannot be empty")
	}

	return nil
}
