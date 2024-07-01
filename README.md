USAGE:

- to generate password of length 15 `go run main.go 15`


Build app `go build -o password-generator password-generator.go`

now you can execute it like this `go build -o password-generator password-generator.go`

install goreleaser and generate a new release
```sh
go get -u github.com/goreleaser/goreleaser
export GITHUB_TOKEN="your github token"
git tag v1.0.0
goreleaser release
```