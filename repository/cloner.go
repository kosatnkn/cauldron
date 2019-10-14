package repository

import (
	"fmt"
	"os"

	"github.com/kosatnkn/cauldron/log"
	"gopkg.in/src-d/go-git.v4"
)

const catalystURL string = "https://github.com/kosatnkn/catalyst.git"

// Clone clones `Catalyst` to the given directory.
func Clone(dir string) (*git.Repository, error) {

	m := fmt.Sprintf("\nCloning Catalyst from %s", catalystURL)
	log.Info(m)

	opts := &git.CloneOptions{
		URL:      catalystURL,
		Progress: os.Stdout,
	}

	r, err := git.PlainClone(dir, false, opts)
	if err != nil {
		return nil, err
	}

	return r, nil
}
