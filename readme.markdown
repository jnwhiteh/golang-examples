## Example project for Go

This project consists of a fairly simple example setup for writing programs
using the Go programming language. It can produce two packages complete with
tests, and two commands, one of which depends on the package in the repository.

The `go` tool takes care of building dependencies for us. To start, you'll need
to set the `GOPATH` environment variable, which can be done by running
`envsetup.sh` like so:

     . envsetup.sh

The two commands can be built using `go install hello` and `go install buffer`,
or you can install them both using `go install hello buffer`. To run the tests
for the packages, you can run `go test jnwhiteh.net/buffers` and `go test jnwhiteh.net/merge`.
