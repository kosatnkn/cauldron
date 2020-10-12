package repository

import (
	"github.com/kosatnkn/cauldron/log"
	"gopkg.in/src-d/go-git.v4"
)

// Init initializes a new repository.
func Init(dir string) error {

	log.Info("Initializing as a new git repo")
	log.Note(" You will have to manually run `git add .` to stage everything in the newly created project")

	_, err := git.PlainInit(dir, false)

	return err
}
