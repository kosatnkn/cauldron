package project

import (
	"fmt"
	"os"
	"regexp"
)

// considered files
var consideredFiles = []string{
	".go",
	"go.mod",
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
