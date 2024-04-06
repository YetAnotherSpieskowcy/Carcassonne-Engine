.PHONY: build
build:
	go build -o bin/ ./...

.PHONY: run
run: build
	./bin/main

.PHONY: test
test:
	go test ./...
