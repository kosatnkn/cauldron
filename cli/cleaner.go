package cli

import (
	"strings"

	"github.com/kosatnkn/cauldron/config"
)

// CleanConfig cleans values in config
func CleanConfig(cfg *config.Config) {

	cfg.Project.Name = cleanString(cfg.Project.Name)
	cfg.Project.Namespace = cleanString(cfg.Project.Namespace)
}

// clean cleans the input value.
func cleanString(value string) string {

	// remove spaces
	value = strings.Replace(value, " ", "", -1)

	return value
}
