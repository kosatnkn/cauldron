package main

import (
	"github.com/kosatnkn/cauldron/repository"
)

func main() {

	name := "sample"
	version := "v1.0.0"

	// clone and checkout tag
	repository.GetBase(name, version)

	// reconfigure project
	// project.Reconfigure("sample")
}
