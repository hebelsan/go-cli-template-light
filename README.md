# go lightweight cobra template

## Installation
```bash
# writes binary to GOPATH
make build
go-template-cli
```
or
```bash
go run main.go
```

## Release
```bash
# dry run
goreleaser release --snapshot --clean
# verify your .goreleaser.yaml
goreleaser check
# do release
export GITHUB_TOKEN="YOUR_GH_TOKEN"
git tag -a v0.1.0 -m "First release"
git push origin v0.1.0
goreleaser release
```

## Customize template
- adjust module name
- go through all files and rename go-template-cli with your project name
- adjust GO_VERSION in Dockerfile
- go version in go.mod
- for each new command add a file in cmd/
- adjust repo name in .goreleaser (replace "github.com/hebelsan/go-template-cli" with your repo)

### Generate Markdown Documentation
```
make doc
```

### Using cobra cli

#### install
go install github.com/spf13/cobra-cli@latest

#### create new command
cobra-cli add <CMD_NAME>