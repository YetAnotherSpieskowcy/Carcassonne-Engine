# Carcassonne-Engine

A rule engine for the Carcassonne game.

## Pre-requirements

### Linux

1. Install Go 1.22 either from your distro's package repositories or by following [instructions on Golang's site](https://go.dev/doc/install)

   **Tip:** If you're using Ubuntu 23.10 or lower, Go version in official repositories is going to be too old.
   You can get the latest version by adding [the PPA listed on Go wiki](https://go.dev/wiki/Ubuntu) and installing `golang` package after.
2. Install `gcc` toolchain from your distro's package repositories (for example, `build-essential` package on Ubuntu).
3. Install Python 3.12 (default version on Ubuntu 24.04, find how to install on your distribution otherwise)
4. Install [Docker Engine](https://docs.docker.com/engine/install/).

### Windows

> [!NOTE]
> Windows support is maintained on best-effort basis - we may drop it once it becomes a burden.

We provide you with 2 ways of installing pre-requirements - manual installation or automated installation with Chocolatey:

#### Manually installing pre-requirements

1. Install [Go for Windows x86-64](https://go.dev/dl/)
2. Install [MinGW-w64](https://github.com/niXman/mingw-builds-binaries)

   You will have to add following path to PATH in your **User** Environment Variables after installation:
   ```console
   %USERPROFILE%\mingw64\bin
   ```
3. Install [Python 3.12](https://www.python.org/downloads/windows/).
4. Install [Docker Desktop (or Engine)](https://docs.docker.com/desktop/install/windows-install/)

> [!NOTE]
> This will require enabling optional Windows features (such as Hyper-V) and may require a reboot.

#### Installing pre-requirements with Chocolatey

*These instructions assume that you already have Chocolatey installed. If not, you can install it by following [its install documentation](https://chocolatey.org/install).*

> [!NOTE]
> The instructions below will enable the Hyper-V feature on your system.

Run PowerShell as Administrator and execute the following command:
```console
choco install Containers Microsoft-Hyper-V --source windowsfeatures
choco install docker-engine golang mingw golangci-lint python312
```

Reboot the system to enable Hyper-V and start the Docker service.

## Building sources

You can either use the default make target:
- Linux
```console
make
```
- Windows
```console
Set-ExecutionPolicy Bypass -Scope Process -Force
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
Set-ExecutionPolicy Bypass -Scope Process -Force
./make.ps1 test
```
or run the test command manually:
```console
go test -race "-coverprofile=coverage.txt" "./..."
```

To show coverage, you can either use the `open-coverage` make target:
- Linux
```console
make open-coverage
```
- Windows
```console
Set-ExecutionPolicy Bypass -Scope Process -Force
./make.ps1 open-coverage
```
or run the cover tool command manually:
```console
go tool cover "-html=coverage.txt"
```

## Linting

You can either use the `lint` make target:
- Linux
```console
make lint
```
- Windows
```console
Set-ExecutionPolicy Bypass -Scope Process -Force
./make.ps1 lint
```
or run the lint command manually:
```console
docker run -e "VALIDATE_ALL_CODEBASE=true" -e "DEFAULT_BRANCH=origin/main" -e "VALIDATE_GO=false" -e "LOG_LEVEL=NOTICE" -e "RUN_LOCAL=true" -v ".:/tmp/lint" --rm "ghcr.io/super-linter/super-linter:v6.3.1"
```


## Performance test
To run the performance tests, go needs to be built with build tag.
Build tags in visual studio code can be set in .vscode/setting.json,
which looks like this:
{    
    "go.buildTags": "test,performanceTest"
}
("test" tag is for end tests)
Tag is made to prevent running long tests on github