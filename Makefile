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
	docker run -e "VALIDATE_ALL_CODEBASE=true" -e "DEFAULT_BRANCH=origin/main" -e "VALIDATE_GO=false" -e "LOG_LEVEL=NOTICE" -e "RUN_LOCAL=true" -v ".:/tmp/lint" --rm "ghcr.io/super-linter/super-linter:latest"

.PHONY: open-coverage
open-coverage:
	go tool cover "-html=coverage.txt"
	@echo "Coverage opened in the default browser."
