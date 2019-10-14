package project

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/kosatnkn/cauldron/log"
)

// considered file signatures for indexing.
var consideredFiles = []string{
	".go",
	"go.mod",
}

// index indexes all files in the given directory extracting files that matches
// the considered file name patterns list.
func index(dir string) (map[string]string, error) {

	log.Info("\nIndexing")

	idx := make(map[string]string)
	scanned := 0
	considered := 0

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		scanned++

		if isConsidered(info) {

			considered++
			idx[info.Name()] = path
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	m := fmt.Sprintf("Indexed %d matching files after scanning %d entries", considered, scanned)
	log.Default(m)

	return idx, nil
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
