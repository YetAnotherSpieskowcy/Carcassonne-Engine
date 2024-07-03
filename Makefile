.PHONY: build
build: .venv
	@echo "Building the project..."
	go build "./pkg/..."
	.venv/bin/python -m pip install .

.PHONY: test
test:
	@echo "Running the test suite..."
	go test -race "-coverprofile=coverage.txt" "./pkg/..."

.PHONY: lint
lint:
	@echo "Running the linter..."
	docker run -e "VALIDATE_ALL_CODEBASE=true" -e "DEFAULT_BRANCH=origin/main" -e "VALIDATE_GO=false" -e "FILTER_REGEX_EXCLUDE=.*python_bindings/.*" -e "VALIDATE_PYTHON_PYLINT=false" -e "LOG_LEVEL=NOTICE" -e "RUN_LOCAL=true" -v ".:/tmp/lint" --mount "type=tmpfs,destination=/tmp/lint/python_bindings" --rm "ghcr.io/super-linter/super-linter:v6.3.1"

.PHONY: open-coverage
open-coverage:
	go tool cover "-html=coverage.txt"
	@echo "Coverage opened in the default browser."

.venv:
	python3.12 -m venv .venv
