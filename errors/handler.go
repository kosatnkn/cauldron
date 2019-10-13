package errors

import (
	"os"

	"github.com/kosatnkn/cauldron/log"
)

// Handle gracefully handles an error.
func Handle(err error) {

	log.Error(err.Error())
	log.Info("Cauldron stopped...")

	os.Exit(1)
}
