package config

// Config holds all configurations.
type Config struct {
	Cauldron Cauldron
	Project  Project
	Base     Base
}

// Cauldron holds configurations of the app.
type Cauldron struct {
	Version string
}

// Project holds new project configurations.
type Project struct {
	Name        string
	Namespace   string
	SplashStyle string
}

// Base holds configurations of the base project
// that will be used to create a new project.
type Base struct {
	Repo        string
	Module      string
	Version     string
	MinVersion  string
	MaxVersion  string
	RemoveDirs  []string
	RemoveFiles []string
}
