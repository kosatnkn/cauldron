package project

import "fmt"

func Reconfigure(name string) {

	createSplash(name)
}

func createSplash(name string) {

	fmt.Printf("Creating splash for %s \n", name)
}
