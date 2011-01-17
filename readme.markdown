## Example project for Go

This is a very simple example project for the Go programming language that
includes two executable commands, a single package (library in Go) and a
top-level makefile that helps tie things together a bit. One of the commands is
standalone, while the other imports from the package that is being built.

The package is installed in your $GOROOT using the 'make install' target.

If you do not have the $GOROOT environment variable set, then you should ensure
you use the `gomake` program instead of the standard `make` program, as it
ensures the correct variables are set.
