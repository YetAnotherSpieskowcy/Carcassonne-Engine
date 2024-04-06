# Carcassonne-Engine

A rule engine for the Carcassonne game.

## Building sources

You can either use the default make target:
```bash
make
```
or run the build command manually:
```bash
go build -o bin/ ./...
```

This will build all Go source files *and* output all built binaries to the `bin/` folder.

## Running the engine

You can either use the `run` make target:
```bash
make run
```
or run the binary in the **Building sources** section:
```bash
./bin/main
```

## Running the test suite

You can either use the `test` make target:
```bash
make test
```
or run the binary in the **Building sources** section:
```bash
go test ./...
```
