package project

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/kosatnkn/cauldron/content"
)

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

	splash := []byte(content.GenerateSplashStyle(name))

	err := ioutil.WriteFile(file, splash, 0666)
	if err != nil {
		return err
	}

	return nil
}
