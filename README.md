# Cauldron

[![Actions Status](https://github.com/kosatnkn/cauldron/workflows/CI/badge.svg)](https://github.com/kosatnkn/cauldron/actions)
[![Coverage Status](https://coveralls.io/repos/github/kosatnkn/cauldron/badge.svg?branch=master)](https://coveralls.io/github/kosatnkn/cauldron?branch=master)
![Open Issues](https://img.shields.io/github/issues/kosatnkn/cauldron)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/kosatnkn/cauldron)
[![Go Reference](https://pkg.go.dev/badge/github.com/kosatnkn/cauldron/v2.svg)](https://pkg.go.dev/github.com/kosatnkn/cauldron/v2)
[![Go Report Card](https://goreportcard.com/badge/github.com/kosatnkn/cauldron)](https://goreportcard.com/report/github.com/kosatnkn/cauldron)

Project generator using [Catalyst](https://github.com/kosatnkn/catalyst) as a template.

`Cauldron` will generate a fully configured project along with a sample set so you can check whether everything is working fine.

The sample set covers request response lifecycle handling of basic CRUD operations. So you can use them as an example to create your REST API.

The project that will be created uses `go.mod` for dependency management. This will enable you to create and run projects outside of the `GOPATH`.

Visit the [Catalyst](https://github.com/kosatnkn/catalyst) base project for more information.

## Version 1 vs Version 2
Version 1 of `Cauldron` supports creating projects using `Catalyst` versions `v1.0.0` to `v2.3.x`.

Version 2 of `Cauldron` supports creating projects using `Catalyst` version `v2.4.0` and upwards.

## Installation

**Version 1**
```bash
go get github.com/kosatnkn/cauldron
```

**Version 2**
```bash
go get github.com/kosatnkn/cauldron/v2
```

## Usage

**Command**
```bash
cauldron -n Sample -s github.com/username [-t v1.0.0]
```
```bash
cauldron --name Sample --namespace github.com/username [--tag v1.0.0]
```

**Input Parameters**
- `-n --name` Project name (ex: ProjectOne). The name will be converted to lowercase to be used in module path.
- `-s --namespace` Namespace for the project (ex: github.com/example)
- `-t --tag` Release version of `Catalyst` to be used. The latest version will be used if `-t` is not provided
- `-h --help` Show help message

This will create a new project with **go.mod** module path of `github.com/username/sample`

Cauldron will do a `git init` on the newly created project but you will have to stage all the files in the project and do the first commit yourself.
```bash
git add .

git commit -m "first commit"
```
