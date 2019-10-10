package project

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

// considered files
var consideredFiles = []string{
	".go",
	"go.mod",
}

var currentModule = "github.com/kosatnkn/catalyst"

// configure configures the project.
func configure(name string, module string, dir string) error {

	configureImportPaths(dir, currentModule, module)

	// configureSplash(name)

	return nil
}

// configureImportPaths rewrites all Catalyst import path segments
// with the given import path segment.
func configureImportPaths(dir string, currentModule string, module string) error {

	fmt.Println("Configuring imports")

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if isConsidered(info) {

			fmt.Println("Name:", info.Name(), "Path:", path)

			replaceContent(path, currentModule, module)
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// // configureSplash configures the splash message.
// func configureSplash(name string) error {

// 	fmt.Println("Creating splash for", name)

// 	return nil
// }

// replaceContent replaces matching strings in file with the given string.
func replaceContent(file string, phrase string, replacement string) error {

	current, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	output := bytes.Replace(current, []byte(phrase), []byte(replacement), -1)

	err = ioutil.WriteFile(file, output, 0666)
	if err != nil {
		return err
	}

	return nil
}

// isConsidered checks whether the directory or the file is ignored.
func isConsidered(info os.FileInfo) bool {

	if info.IsDir() {
		return false
	}

	for _, file := range consideredFiles {

		r, err := regexp.Compile(file)
		if err != nil {
			fmt.Println(err)
		}

		if r.MatchString(info.Name()) {
			return true
		}
	}

	return false
}
