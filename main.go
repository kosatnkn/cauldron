package main

import (
	"github.com/kosatnkn/cauldron/cli"
	"github.com/kosatnkn/cauldron/config"
	"github.com/kosatnkn/cauldron/project"
)

func main() {

	cli.ShowSplash()

	// read input from stdin

	// set config
	cfg := &config.Config{
		Name:        "Sample",
		Namespace:   "test.com/sampleuser",
		Tag:         "",
		SplashStyle: "shadow",
	}

	// validate config
	config.Validate(cfg)

	// create project
	project.Create(cfg)
}
