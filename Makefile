.PHONY: build
build:
	@echo "Building the project..."
	go build "./..."

.PHONY: test
test:
	@echo "Running the test suite..."
	go test -race "-coverprofile=coverage.txt" "./..."

.PHONY: open-coverage
open-coverage:
	go tool cover "-html=coverage.txt"
	@echo "Coverage opened in the default browser."
