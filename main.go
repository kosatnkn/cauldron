package main

import (
	"github.com/kosatnkn/cauldron/v2/cli"
	"github.com/kosatnkn/cauldron/v2/config"
	"github.com/kosatnkn/cauldron/v2/project"
)

func main() {

	// init config
	cfg := &config.Config{
		Cauldron: config.Cauldron{
			Version: "v2.0.0",
		},
		Project: config.Project{
			SplashStyle: "small",
		},
		Base: config.Base{
			Repo:       "https://github.com/kosatnkn/catalyst.git",
			Module:     "github.com/kosatnkn/catalyst",
			MinVersion: "v1.0.0",
			MaxVersion: "v2.3.0",
			RemoveDirs: []string{
				".git",
				".github",
			},
			RemoveFiles: []string{
				".travis.yml",
				"doc.go",
			},
		},
	}

	cli.ParseFlags(cfg)

	cli.ShowSplash(cfg)

	// clean config
	cli.CleanConfig(cfg)

	// print config values
	cli.ShowConfig(cfg)

	// validate config
	config.Validate(cfg)

	// create project
	project.Create(cfg)

	// complete
	cli.ShowComplete(cfg)
}
