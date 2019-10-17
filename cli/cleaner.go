package cli

import (
	"strings"

	"github.com/kosatnkn/cauldron/config"
)

// CleanConfig cleans values in config
func CleanConfig(cfg *config.Config) {

	cfg.Name = cleanString(cfg.Name)
	cfg.Namespace = cleanString(cfg.Namespace)
}

// clean cleans the input value.
func cleanString(value string) string {

	// remove spaces
	value = strings.Replace(value, " ", "", -1)

	return value
}
