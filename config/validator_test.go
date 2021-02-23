package config

import "testing"

func TestCheckNameValid(t *testing.T) {

	tests := []struct {
		name  string
		input string
	}{
		{"single word", "Sample"},
		{"two words", "Sample One"},
		{"many words", "Sample with many words"},
		{"underscores", "Sample_with underscores"},
		{"dashes", "Sample-with dashes"},
		{"slashes", "Sample/with slashes"},
	}
	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			err := checkNameValid(test.input)
			if err != nil {
				t.Errorf("Need nil, got %v", err)
			}
		})
	}
}

func TestCheckNameValidFail(t *testing.T) {

	err := checkNameValid("")

	if err == nil {
		t.Fatalf("Need error, got nil")
	}

	need := "Project name cannot be empty"
	got := err.Error()

	if got != need {
		t.Errorf("Need %s, got %s", need, got)
	}
}

func TestCheckNameSpaceValid(t *testing.T) {

	tests := []struct {
		name  string
		input string
	}{
		{"single word", "Sample"},
		{"slashes", "Sample/with/slashes"},
	}
	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			err := checkNameSpaceValid(test.input)
			if err != nil {
				t.Errorf("Need nil, got %s", err.Error())
			}
		})
	}
}

func TestCheckNameSpaceValidFail(t *testing.T) {

	tests := []struct {
		name  string
		input string
		need  string
	}{
		{"single word", "", "NameSpace cannot be empty"},
		{"two words", "Sample One", "Namespace contain illegal characters"},
		{"many words", "Sample with many words", "Namespace contain illegal characters"},
		{"underscores", "Sample_with underscores", "Namespace contain illegal characters"},
		{"dashes", "Sample-with dashes", "Namespace contain illegal characters"},
		{"slashes", "Sample/with slashes", "Namespace contain illegal characters"},
		{"special characters", `Sample@_#with% \ /special chars`, "Namespace contain illegal characters"},
	}
	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			err := checkNameSpaceValid(test.input)
			if err.Error() != test.need {
				t.Errorf("Need %s, got %s", test.need, err.Error())
			}
		})
	}
}

func TestCheckValidSemanticTag(t *testing.T) {

	tests := []struct {
		name  string
		input string
	}{
		{"symantic major", "v1.0.0"},
		{"symantic minor", "v1.2.0"},
		{"symantic patch", "v1.2.3"},
		{"double digits", "v11.22.33"},
		{"pre release", "v1.2.3-a1"},
	}
	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			err := checkValidSemanticTag(test.input)
			if err != nil {
				t.Errorf("Need nil, got %v", err)
			}
		})
	}
}

func TestCheckValidSemanticTagFail(t *testing.T) {

	tests := []struct {
		name  string
		input string
		need  string
	}{
		{"one point", "v1", `'v1' is not a valid semantic tag`},
		{"two points", "v1.2", `'v1.2' is not a valid semantic tag`},
		{"without v", "1.2.3", `'1.2.3' is not a valid semantic tag`},
	}
	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			err := checkValidSemanticTag(test.input)
			if err == nil {
				t.Errorf("Need error, got nil")
			}

			got := err.Error()
			if got != test.need {
				t.Errorf("Need %s, got %s", test.need, got)
			}
		})
	}
}
