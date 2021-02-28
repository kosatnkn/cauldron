package config

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/Masterminds/semver/v3"
	e "github.com/kosatnkn/cauldron/v2/errors"
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

	err = checkValidSemanticTag(cfg.Base.Version)
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

// checkValidSemanticTag checks whether the provided tag is a valid semantic tag
func checkValidSemanticTag(tag string) error {

	if tag == "" {
		return nil
	}

	exp := `^v(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)`

	if !regexp.MustCompile(exp).MatchString(tag) {
		return fmt.Errorf(`'%s' is not a valid semantic tag`, tag)
	}

	return nil
}

// checkVersionInRange checks whether the provided version is in the supported version range.
func checkVersionInRange(cfg *Config) error {

	if cfg.Base.Version == "" {
		return nil
	}

	v, err := semver.NewVersion(cfg.Base.Version)
	if err != nil {
		return err
	}

	min := cfg.Base.MinVersion[1:]
	max := cfg.Base.MaxVersion[1:]
	c, err := semver.NewConstraint(fmt.Sprintf(">=%s, <=%s", min, max))
	if err != nil {
		return err
	}

	ok, _ := c.Validate(v)
	if !ok {
		return fmt.Errorf("Version out of range")
	}

	return nil
}
