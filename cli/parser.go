package cli

import (
	"flag"

	"github.com/kosatnkn/cauldron/config"
)

// ParseFlags parses input flags and populates the config object.
func ParseFlags(cfg *config.Config) {

	// read input from stdin
	flag.StringVar(&cfg.Project.Name,
		"n",
		"",
		"Name for the new project.")

	flag.StringVar(&cfg.Project.Namespace,
		"ns",
		"",
		"Namespace for the new project.")

	flag.StringVar(&cfg.Base.Version,
		"t",
		"",
		`Release tag of the base project that is going to be used to create the new project.
When a tag is not defined the latest tag will be used`)

	flag.Parse()
}
