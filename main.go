package main

import (
	"github.com/kosatnkn/cauldron/cli"
	"github.com/kosatnkn/cauldron/project"
)

func main() {

	cli.ShowSplash()

	// read input from stdin

	name := "sample"
	version := "v1.1.0"
	modulePart := "test.com/sampleuser"

	// create project
	project.Create(name, modulePart, version)
}
