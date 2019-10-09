package main

import "github.com/kosatnkn/cauldron/repository"

func main() {

	// command
	// $ cauldron

	// get input params
	// - New project name (have an input tag --name / -n)
	// - Catalyst version to be used (have an input tag --tag=v1.0.0 / -t (default: latest))
	// - openapi.yaml file (have an input tag --spec)
	//  - validate naming conventions (have an input tag --strict=false / -s (default: true) to ignore this validaton)
	//  - create base structures for unpackers using request body
	//  - create transformers using response body

	// interactive mode
	// when tags are not provided use interactive mode

	name := "sample"
	version := "v1.0.0"

	// process
	// - clone
	// - checkout tag
	repository.GetBase(name, version)

	// - remove .git
	// project.Clean()

	// - generate structures
	// - create splash using name
	// project.Update("sample")

	// - init as new git repo
	// repository.Init()
}
