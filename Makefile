.PHONY: build
build:
	@echo "Building the project..."
	go build "./..."

.PHONY: test
test:
	@echo "Running the test suite..."
	go test -race "-coverprofile=coverage.txt" "./..."

.PHONY: lint
lint:
	@echo "Running the linter..."
	golangci-lint run

.PHONY: open-coverage
open-coverage:
	go tool cover "-html=coverage.txt"
	@echo "Coverage opened in the default browser."
