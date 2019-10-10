package project

import (
	"fmt"
	"os"
)

// directories to be removed
var dirs = []string{
	".git",
}

// files to be removed
var files = []string{
	"doc.go",
	"README.md",
}

// clean cleans the project directory.
func clean(baseDir string) error {

	fmt.Println("Cleaning")

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

		fmt.Println("Removed", dir)
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

		fmt.Println("Removed", file)
	}

	return nil
}
