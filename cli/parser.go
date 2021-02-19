package cli

import (
	"github.com/spf13/pflag"

	"github.com/kosatnkn/cauldron/config"
)

// ParseFlags parses input flags and populates the config object.
func ParseFlags(cfg *config.Config) {

	// read input from stdin
	pflag.StringVarP(&cfg.Project.Name,
		"name",
		"n",
		"",
		`Name for the new project
Only alpha-numeric characters dash(-) and underscore(_) are allowed. Any other character will be removed.
e.g.: Sample
`)

	pflag.StringVarP(&cfg.Project.Namespace,
		"namespace",
		"s",
		"",
		`Namespace for the new project
Only alpha-numeric characters dash(-), underscore(_) and slash(/) are allowed. Any other character will be removed.
e.g.: github.com/username
`)

	pflag.StringVarP(&cfg.Base.Version,
		"tag",
		"t",
		"",
		`Release version of the base project to be used (Optional)
When a tag is not defined the latest tag will be used.
e.g.: v1.0.0
`)

	pflag.Parse()
}
