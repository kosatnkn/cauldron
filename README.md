# Cauldron

Project generator using [Catalyst](https://github.com/kosatnkn/catalyst) as a template.

`Cauldron` will generate a fully configured project along with a sample set so you can check whether everything is working fine.

The sample set covers request response lifecycle handling of basic CRUD operations. So you can use them as an example to create your REST API.

The project that will be created uses `go.mod` for dependency management. This will enable you to create and run projects outside of the `GOPATH`.

Visit the [Catalyst](https://github.com/kosatnkn/catalyst) base project for more information.

## Installation

**Clone Cauldron**
```bash
    git clone https://github.com/kosatnkn/cauldron.git
```

**Install**
```bash
    cd cauldron

    go install
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
