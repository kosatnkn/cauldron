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
		Namespace:   "",       // test.com/sampleuser
		Tag:         "",       // latest
		SplashStyle: "shadow", // shadow
	}

	cli.ParseFlags(cfg)

	cli.ShowSplash()

	// print config values
	cli.ShowConfig(cfg)

	// validate config
	config.Validate(cfg)

	// create project
	project.Create(cfg)

	// git init
}
