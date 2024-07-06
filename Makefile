.PHONY: build-and-install
build-and-install: build-go install-python

.PHONY: build
build: build-go build-python

.PHONY: build-go
build-go:
	@echo "Building the Go project..."
	go build "./pkg/..."

.PHONY: build-python
build-python: .venv
	@echo "Generating and building Python bindings..."
	@mkdir -p built_wheels
	go install "github.com/go-python/gopy@v0.4.10"
	go install "golang.org/x/tools/cmd/goimports@latest"
	.venv/bin/python -m pip wheel . --wheel-dir=built_wheels

.PHONY: install-python
install-python: build-python
	.venv/bin/python -m pip install carcassonne_engine \
		--no-index --find-links=built_wheels \
		--force-reinstall -U

.PHONY: test
test: test-go test-python

.PHONY: test-go
test-go:
	@echo "Running the Go test suite..."
	go test -race "-coverprofile=coverage.txt" "./pkg/..."

.PHONY: test-python
test-python: install-python
	@echo "Running the Python test suite..."
	.venv/bin/python -m pytest -s

.PHONY: lint
lint:
	@echo "Running the linter..."
	docker run \
		-e "VALIDATE_ALL_CODEBASE=true" \
		-e "DEFAULT_BRANCH=origin/main" \
		-e "VALIDATE_GO=false" \
		-e "FILTER_REGEX_EXCLUDE=.*python_bindings/.*" \
		-e "VALIDATE_PYTHON_PYLINT=false" \
		-e "LOG_LEVEL=NOTICE" \
		-e "RUN_LOCAL=true" \
		-v ".:/tmp/lint" \
		--mount "type=tmpfs,destination=/tmp/lint/python_bindings" \
		--rm \
		"ghcr.io/super-linter/super-linter:v6.3.1"

.PHONY: open-coverage
open-coverage:
	go tool cover "-html=coverage.txt"
	@echo "Coverage opened in the default browser."

.venv:
	python3.12 -m venv .venv
	.venv/bin/python -m pip install -r requirements-dev.txt
