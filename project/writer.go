package project

import (
	"bytes"
	"io/ioutil"
)

// replaceContent replaces matching strings in file with the given string.
func replaceContent(file string, phrase string, replacement string) error {

	current, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	output := bytes.Replace(current, []byte(phrase), []byte(replacement), -1)

	return write(file, output)
}

// write writes the given content to the specified file.
func write(file string, content []byte) error {

	err := ioutil.WriteFile(file, content, 0666)
	if err != nil {
		return err
	}

	return nil
}
