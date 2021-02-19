package cli

import (
	"testing"
)

func TestCleanString(t *testing.T) {

	tests := []struct {
		name  string
		input string
		need  string
	}{
		{"single word", "Sample", "Sample"},
		{"two words", "Sample One", "SampleOne"},
		{"many words", "Sample with many words", "Samplewithmanywords"},
		{"underscores", "Sample_with underscores", "Sample_withunderscores"},
		{"dashes", "Sample-with dashes", "Sample-withdashes"},
		{"slashes", "Sample/with slashes", "Sample/withslashes"},
		{"special characters", `Sample@_#with% \ /special chars`, "Sample_with/specialchars"},
	}
	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			got := cleanString(test.input)
			if got != test.need {
				t.Errorf("Need %s, got %s", test.need, got)
			}
		})
	}
}
