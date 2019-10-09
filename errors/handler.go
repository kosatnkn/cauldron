package errors

import (
	"fmt"
	"os"
)

// Handle gracefully handles an error.
func Handle(err error) {

	fmt.Println("Error:", err)
	fmt.Println("Cauldron stopped")

	os.Exit(1)
}
