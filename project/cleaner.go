package project

import (
	"fmt"
	"os"

	"github.com/kosatnkn/cauldron/log"
)

// directories to be removed
var dirs = []string{
	".git",
	".github",
}

// files to be removed
var files = []string{
	"doc.go",
}

// clean cleans the project directory.
func clean(baseDir string) error {

	log.Info("Cleaning")

	baseDir = baseDir + string(os.PathSeparator)

	// delete directories
	err := deleteDirs(baseDir, dirs)
	if err != nil {
		return err
	}

	// delete files
	err = deleteFiles(baseDir, files)
	if err != nil {
		return err
	}

	return nil
}

// deleteDir deletes the given directories.
func deleteDirs(baseDir string, dirs []string) error {

	for _, dir := range dirs {

		err := os.RemoveAll(baseDir + dir)
		if err != nil {
			return err
		}

		m := fmt.Sprintf(" Removed %s", dir)
		log.Default(m)
	}

	return nil
}

// deleteFile deletes the given files.
func deleteFiles(baseDir string, files []string) error {

	for _, file := range files {

		err := os.Remove(baseDir + file)
		if err != nil {
			return err
		}

		m := fmt.Sprintf(" Removed %s", file)
		log.Default(m)
	}

	return nil
}
