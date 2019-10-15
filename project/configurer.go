package project

import (
	"os"
	"strings"

	"github.com/kosatnkn/cauldron/content"
	"github.com/kosatnkn/cauldron/log"
)

var currentModule = "github.com/kosatnkn/catalyst"

// configure configures the project.
func configure(name string, simpleName string, modulePrefix string, files map[string]string) error {

	log.Info("\nConfiguring")

	module := getModule(modulePrefix, simpleName)
	var err error

	for k, file := range files {

		// NOTE: allow to fall through without jumping to the next iteration
		// so that import paths will be properly rewritten

		// rewrite splash message
		if k == "styles.go" {
			err = rewriteSplash(file, name)
			if err != nil {
				return err
			}
		}

		// rewrite readme
		if k == "README.md" && isBaseReadme(file, simpleName) {
			err = rewriteReadme(file, name)
			if err != nil {
				return err
			}
		}

		// rewrite import paths
		err = rewriteImportPaths(file, module)
		if err != nil {
			return err
		}
	}

	return nil
}

// isBaseReadme checks whether the given readme file is the main readme
// file at the base of the project.
func isBaseReadme(file string, simpleName string) bool {

	parts := strings.Split(file, string(os.PathSeparator))

	// checks to see whether the element right before the last is
	// the name of the base directory
	return parts[len(parts)-2] == simpleName
}

// rewriteImportPaths replaces the old module name in the
// import path with the new module name.
func rewriteImportPaths(file string, module string) error {

	log.Default(file)

	return replaceContent(file, currentModule, module)
}

// rewriteSplash creates a new splash for the project using the project name.
func rewriteSplash(file string, name string) error {

	splash := []byte(content.GenerateSplashStyle(name))

	return write(file, splash)
}

// rewriteReadme creates a new readme file for the project.
func rewriteReadme(file string, name string) error {

	readme := []byte(content.GenerateReadme(name))

	return write(file, readme)
}
