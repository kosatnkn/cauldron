package repository

import (
	"fmt"
	"os"

	"github.com/kosatnkn/cauldron/v2/log"
	"gopkg.in/src-d/go-git.v4"
)

// Clone clones `Catalyst` to the given directory.
func Clone(dir string, url string) (*git.Repository, error) {

	m := fmt.Sprintf("Cloning %s", url)
	log.Info(m)

	opts := &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	}

	r, err := git.PlainClone(dir, false, opts)
	if err != nil {
		return nil, err
	}

	return r, nil
}
