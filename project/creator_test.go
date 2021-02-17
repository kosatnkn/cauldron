package project

import (
	"testing"
)

func TestVersionInRange(t *testing.T) {

	b := isVersionInRange("v2.3.0", "v1.0.0", "v2.3.0")

	if !b {
		t.Errorf("Validation failed")
	}
}
