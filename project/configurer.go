package project

import "fmt"

// configure configures the project.
func configure(name string, dir string) error {

	configureImports(dir)

	configureSplash(name)

	return nil
}

func configureImports(dir string) error {

	fmt.Println("Configuring imports")

	return nil
}

// configureSplash configures the splash message.
func configureSplash(name string) error {

	fmt.Println("Creating splash for", name)

	return nil
}
