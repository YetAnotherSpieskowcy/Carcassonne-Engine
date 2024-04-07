# Carcassonne-Engine

A rule engine for the Carcassonne game.

## Pre-requirements

### Linux

1. Install Go 1.22 either from your distro's package repositories or by following [instructions on Golang's site](https://go.dev/doc/install)

   **Tip:** If you're using Ubuntu 23.10 or lower, Go version in official repositories is going to be too old.
   You can get the latest version by adding [the PPA listed on Go wiki](https://go.dev/wiki/Ubuntu) and installing `golang` package after.
2. Install `gcc` toolchain from your distro's package repositories (for example, `build-essential` package on Ubuntu).

### Windows

We provide you with 2 ways of installing pre-requirements - manual installation or automated installation with Chocolatey:

#### Manually installing pre-requirements

1. Install [Go for Windows x86-64](https://go.dev/dl/)
2. Install [MinGW-w64](https://github.com/niXman/mingw-builds-binaries)

#### Installing pre-requirements with Chocolatey

*These instructions assume that you already have Chocolatey installed. If not, you can install it by following [its install documentation](https://chocolatey.org/install).*

Run PowerShell as Administrator and execute the following command:
```console
choco install golang mingw
```

## Building sources

You can either use the default make target:
- Linux
```console
make
```
- Windows
```console
./make.ps1
```
or run the build command manually:
```console
go build "./..."
```

This will build all Go source files.

## Running the test suite

You can either use the `test` make target:
- Linux
```console
make test
```
- Windows
```console
./make.ps1 test
```
or run the test command manually:
```console
go test -race "-coverprofile=coverage.txt" "./..."
```

To show coverage, you can either use the `show-coverage` make target:
- Linux
```console
make show-coverage
```
- Windows
```console
./make.ps1 show-coverage
```
or run the cover tool command manually:
```console
go tool cover "-html=coverage.txt"
```
