.PHONY: build
build:
	go build ./...

.PHONY: test
test:
	go test -race "-coverprofile=coverage.txt" ./...

.PHONY: open-coverage
open-coverage:
	go tool cover "-html=coverage.txt"
