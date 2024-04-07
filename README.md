# Carcassonne-Engine

A rule engine for the Carcassonne game.

## Pre-requirements

### Linux

1. Install Go 1.22 either from your distro's package repositories or by following these instructions: https://go.dev/doc/install

   **Tip:** If you're using Ubuntu 23.10 or lower, Go version in official repositories is going to be too old.
   You can get the latest version by adding the PPA listed here and installing `golang` package after: https://go.dev/wiki/Ubuntu
2. Install `gcc` toolchain from your distro's package repositories (for example, `build-essential` package on Ubuntu).

## Building sources

You can either use the default make target:
```console
make
```
or run the build command manually:
```console
go build "./..."
```

This will build all Go source files.

## Running the test suite

You can either use the `test` make target:
```console
make test
```
or run the test command manually:
```console
go test -race "-coverprofile=coverage.txt" "./..."
```

To show coverage, you can either use the `show-coverage` make target:
```console
make show-coverage
```
or run the cover tool command manually:
```console
go tool cover "-html=coverage.txt"
```
