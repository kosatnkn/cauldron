package main

import (
	"github.com/kosatnkn/cauldron/cli"
	"github.com/kosatnkn/cauldron/config"
	"github.com/kosatnkn/cauldron/project"
)

func main() {

	// init config
	cfg := &config.Config{
		Name:        "",       // Sample
		Namespace:   "",       // example.com/example
		Tag:         "",       // v1.0.0
		SplashStyle: "shadow", // shadow
		Repo:        "https://github.com/kosatnkn/catalyst.git",
		Version:     "v1.2.1",
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
