package project

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"github.com/kosatnkn/cauldron/content"
)

// considered files
var consideredFiles = []string{
	".go",
	"go.mod",
}

var currentModule = "github.com/kosatnkn/catalyst"

// configure configures the project.
func configure(name string, module string, dir string) error {

	// rename import paths
	err := configureImportPaths(name, dir, currentModule, module)
	if err != nil {
		return err
	}

	return nil
}

// configureImportPaths rewrites all Catalyst import path segments
// with the given import path segment.
func configureImportPaths(name string, dir string, currentModule string, module string) error {

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if isConsidered(info) {

			fmt.Println("Name:", info.Name(), "Path:", path)

			replaceContent(path, currentModule, module)
		}

		if info.Name() == "styles.go" {

			// create splash
			configureSplash(name, path)
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// configureSplash configures the splash message.
func configureSplash(name string, file string) error {

	fmt.Println("Creating splash for", name)

	splash := []byte(content.GenerateSplash(name))

	err := ioutil.WriteFile(file, splash, 0666)
	if err != nil {
		return err
	}

	return nil
}

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
