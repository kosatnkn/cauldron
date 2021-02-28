package project

import (
	"fmt"
	"os"

	"github.com/kosatnkn/cauldron/log"
)

// clean cleans the project directory.
func clean(baseDir string, dirs, files []string) error {

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

		info, _ := os.Stat(baseDir + file)
		if info == nil {
			log.Warn(fmt.Sprintf(" Missing %s: %s%s", file, baseDir, file))
			continue
		}

		err := os.Remove(baseDir + file)
		if err != nil {
			return err
		}

		m := fmt.Sprintf(" Removed %s", file)
		log.Default(m)
	}

	return nil
}
