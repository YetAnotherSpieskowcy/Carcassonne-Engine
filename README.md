# Carcassonne-Engine

A rule engine for the Carcassonne game.

## Building sources

You can either use the default make target:
```bash
make
```
or run the build command manually:
```bash
go build ./...
```

This will build all Go source files.

## Running the test suite

You can either use the `test` make target:
```bash
make test
```
or run the test command manually:
```bash
go test -race "-coverprofile=coverage.txt" ./...
```

To show coverage, you can either use the `show-coverage` make target:
```bash
make show-coverage
```
or run the cover tool command manually:
```bash
go tool cover "-html=coverage.txt"
```
