package project

import (
	"fmt"

	"github.com/kosatnkn/cauldron/content"
	"github.com/kosatnkn/cauldron/log"
)

var currentModule = "github.com/kosatnkn/catalyst"

// configure configures the project.
func configure(name string, modulePrefix string, files map[string]string) error {

	log.Info("\nConfiguring")

	module := getModule(modulePrefix, name)

	for k, file := range files {

		// rewrite splash message
		// NOTE: allow to fall through without jumping to the next iteration
		// so that import paths will be properly rewritten
		if k == "styles.go" {
			rewriteSplash(file, name)
		}

		// rewrite import paths
		err := rewriteImportPaths(file, module)
		if err != nil {
			return err
		}
	}

	return nil
}

// rewriteImportPaths replaces the old module name in the
// import path with the new module name.
func rewriteImportPaths(file string, module string) error {

	m := fmt.Sprintf("Configured: %s", file)
	log.Default(m)

	return replaceContent(file, currentModule, module)
}

// rewriteSplash creates a new splash for the project using the project name.
func rewriteSplash(file string, name string) error {

	splash := []byte(content.GenerateSplashStyle(name))

	return write(file, splash)
}
