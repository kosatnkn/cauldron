# Cauldron

[![Actions Status](https://github.com/kosatnkn/cauldron/workflows/CI/badge.svg)](https://github.com/kosatnkn/cauldron/actions)
![Open Issues](https://img.shields.io/github/issues/kosatnkn/cauldron)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/kosatnkn/cauldron)
[![Go Report Card](https://goreportcard.com/badge/github.com/kosatnkn/cauldron)](https://goreportcard.com/report/github.com/kosatnkn/cauldron)

Project generator using [Catalyst](https://github.com/kosatnkn/catalyst) as a template.

`Cauldron` will generate a fully configured project along with a sample set so you can check whether everything is working fine.

The sample set covers request response lifecycle handling of basic CRUD operations. So you can use them as an example to create your REST API.

The project that will be created uses `go.mod` for dependency management. This will enable you to create and run projects outside of the `GOPATH`.

Visit the [Catalyst](https://github.com/kosatnkn/catalyst) base project for more information.

## Installation

```bash
    $ go get github.com/kosatnkn/cauldron
```

## Usage

**Command**
```bash
    $ cauldron -n=ProjectOne -ns=github.com/example [-t=v1.0.0]
```

**Input Parameters**
- `-n` Project name (ex: ProjectOne)
- `-ns` Namespace for the project (ex: github.com/example)
- `-t` Release version of `Catalyst` to be used. The latest version will be used if -t is not provided
- `-help` or `-h` Show help message

Cauldron will do a `git init` on the newly created project but you will have to stage all the files in the project and do the first commit yourself.
```bash
    git add .

    git commit -m "first commit"
```


## Using Go mod

Go mod is used as the dependency management mechanism. Visit [here](https://github.com/golang/go/wiki/Modules) for more details.

Use go mod in projects that are within the `GOPATH`
```bash
    export GO111MODULE=on
```

Initialize go mod
```bash
    go mod init github.com/my/repo
```

View final versions that will be used in a build for all direct and indirect dependencies
```bash
    go list -m all
```
View available minor and patch upgrades for all direct and indirect dependencies
```bash
    go list -u -m all
```
Update all direct and indirect dependencies to latest minor or patch upgrades (pre-releases are ignored)
```bash
    go get -u or go get -u=patch
```
Build or test all packages in the module when run from the module root directory
```bash
    go build ./... or go test ./...
```
Prune any no-longer-needed dependencies from go.mod and add any dependencies needed for other combinations of OS, architecture, and build tags
```bash
    go mod tidy
```
Optional step to create a vendor directory
```bash
    go mod vendor
```
