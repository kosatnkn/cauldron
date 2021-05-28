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
			Version: "v2.0.1",
		},
		Project: config.Project{
			SplashStyle: "small",
		},
		Base: config.Base{
			Repo:             "https://github.com/kosatnkn/catalyst.git",
			Module:           "github.com/kosatnkn/catalyst/v2",
			MinVersion:       "v2.4.0",
			MaxVersion:       "",
			NextMajorVersion: "v3.0.0",
			RemoveDirs: []string{
				".git",
				".github",
			},
			RemoveFiles: []string{
				"doc.go",
				"LICENSE",
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
