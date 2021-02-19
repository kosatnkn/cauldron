package cli

import (
	"fmt"
	"regexp"

	"github.com/kosatnkn/cauldron/config"
)

// CleanConfig cleans values in config
func CleanConfig(cfg *config.Config) {

	cfg.Project.Name = cleanString(cfg.Project.Name)
	cfg.Project.Namespace = cleanString(cfg.Project.Namespace)
}

// clean cleans the input value.
func cleanString(v string) string {

	exp := "[^a-zA-Z0-9/_-]+"
	reg, err := regexp.Compile(exp)
	if err != nil {
		fmt.Println(err)
	}

	return reg.ReplaceAllString(v, "")
}
