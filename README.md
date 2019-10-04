# Cauldron

Project generator using `Catalyst` as a template.

- Interactive CLI to input project name, Catalyst version and other information to be used by the generator
- Ability to use an `openapi.yaml` document to generate the endpoints, unpackers and transformers and maybe stubs for handler methods
- Do a git init on the generated project


## View GoDoc Locally
```shell
    godoc -http=:6060 -v
```

Navigate to [http://localhost:6060/pkg/github.com/kosatnkn/catalyst/](http://localhost:6060/pkg/github.com/kosatnkn/catalyst/)


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