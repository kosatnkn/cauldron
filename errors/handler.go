package errors

import (
	"os"

	"github.com/kosatnkn/cauldron/v2/log"
)

// Handle gracefully handles an error.
func Handle(err error) {

	log.Error(err.Error())
	log.Warn("Cauldron stopped\n")

	os.Exit(1)
}
