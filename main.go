package main

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/kosatnkn/cauldron/project"
)

func main() {

	myFigure := figure.NewFigure("Catalyst", "shadow", true)
	myFigure.Print()

	// read input from stdin

	name := "sample"
	version := "v1.1.0"
	modulePart := "test.com/sampleuser"

	// create project
	project.Create(name, modulePart, version)
}
