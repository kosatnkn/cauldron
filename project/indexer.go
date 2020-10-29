package project

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"

	"github.com/kosatnkn/cauldron/log"
)

// considered file signatures for indexing.
var consideredFiles = []string{
	".go",
	"README.md",
	"go.mod",
}

// index indexes all files in the given directory extracting files that matches
// the considered file name patterns list.
func index(dir string) (map[string]string, error) {

	log.Info("Indexing")

	idx := make(map[string]string)
	scanned := 0
	considered := 0

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if isConsidered(info) {

			// add a temporary random prefix to the index key to avoid overwriting entries
			// a seperator is added to differentiate the prefix if needed
			k := randomString(10) + "|" + info.Name()
			idx[k] = path

			considered++
		}

		scanned++

		return nil
	})
	if err != nil {
		return nil, err
	}

	m := fmt.Sprintf(" Indexed %d matching files after scanning %d entries", considered, scanned)
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

// randomString creates a random string of given length.
func randomString(n int) string {

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}

	return string(s)
}
