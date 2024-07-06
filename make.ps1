<#
.Synopsis
Makefile script in PowerShell that contains commands useful during development of Carcassonne-Engine.

.Description
Available commands:
   build             Build all Go source files.
   test              Run the test suite.
   open-coverage     Show coverage in the browser after running the test suite.
   lint              Lint all Go source files.

.Parameter Command
Command to execute. See Cmdlet's description for more information.

#>

# I'm too dumb for PowerShell, so $script:availableCommands needs to be defined in 2 places // Jack

[Diagnostics.CodeAnalysis.SuppressMessageAttribute(
    "PSReviewUnusedParameter",
    "",
    Justification = "Parameter is automatically provided by PowerShell which we have no control over."
)]
[CmdletBinding()]
param (
    [Parameter(Mandatory=$false)]
    [ArgumentCompleter({
        param (
            $commandName,
            $parameterName,
            $wordToComplete,
            $commandAst,
            $fakeBoundParameters
        )
        $script:availableCommands = @(
            "build-and-install",
            "build",
            "build-go",
            "build-python",
            "install-python",
            "test",
            "test-go",
            "test-python",
            "lint",
            "open-coverage"
        )
        return $script:availableCommands | Where-Object { $_ -like "$wordToComplete*" }
    })]
    [String]
    $command,
    [switch]
    $help = $false
)

function New-Venv-If-Needed() {
    [Diagnostics.CodeAnalysis.SuppressMessageAttribute(
        "PSUseShouldProcessForStateChangingFunctions",
        "",
        Justification="Not an exported cmdlet"
    )]
    param()

    if (Test-Path .venv) {
        return
    }
    & py -3.12 -m venv .venv
    Exit-On-Fail $LASTEXITCODE
    & .venv\bin\python -m pip install -r requirements-dev.txt
    Exit-On-Fail $LASTEXITCODE
}

function Exit-On-Fail([int]$exitCode) {
    if ($exitCode) {
        Exit $exitCode
    }
}

function build-and-install() {
    build-go
    install-python
}

function build() {
    build-go
    build-python
}

function build-go() {
    Write-Output "Building the Go project..."
    & go build "./pkg/..."
    Exit-On-Fail $LASTEXITCODE
}

function build-python() {
    New-Venv-If-Needed

    Write-Output "Generating and building Python bindings..."
    New-Item -ItemType Directory -Force -Path built_wheels | Out-Null
    & go install "github.com/go-python/gopy@v0.4.10"
    Exit-On-Fail $LASTEXITCODE
    & go install "golang.org/x/tools/cmd/goimports@latest"
    Exit-On-Fail $LASTEXITCODE
    & .venv\Scripts\python.exe -m pip wheel . --wheel-dir=built_wheels
    Exit-On-Fail $LASTEXITCODE
}

function install-python() {
    build-python
    & .venv\Scripts\python.exe -m pip install carcassonne_engine `
        --no-index --find-links=built_wheels `
        --force-reinstall -U
    Exit-On-Fail $LASTEXITCODE
}

function test() {
    test-go
    test-python
}

function test-go() {
    Write-Output "Running the Go test suite..."
    & go test -race "-coverprofile=coverage.txt" "./pkg/..."
    Exit-On-Fail $LASTEXITCODE
}

function test-python() {
    Write-Output "Running the Python test suite..."
    & .venv\Scripts\python.exe -m pytest -s
    Exit-On-Fail $LASTEXITCODE
}

function lint() {
    Write-Output "Running the linter..."
    & docker run --rm `
        -e "VALIDATE_ALL_CODEBASE=true" `
        -e "DEFAULT_BRANCH=origin/main" `
        -e "VALIDATE_GO=false" `
        -e "VALIDATE_PYTHON_PYLINT=false" `
        -e "FILTER_REGEX_EXCLUDE=.*python_bindings/.*" `
        -e "LOG_LEVEL=NOTICE" `
        -e "RUN_LOCAL=true" `
        -v ".:/tmp/lint" `
        --mount "type=tmpfs,destination=/tmp/lint/python_bindings" `
        "ghcr.io/super-linter/super-linter:v6.3.0"
}

function open-coverage() {
    & go tool cover "-html=coverage.txt"
    Write-Output "Coverage opened in the default browser."
}

$script:availableCommands = @(
    "build-and-install",
    "build",
    "build-go",
    "build-python",
    "install-python",
    "test",
    "test-go",
    "test-python",
    "lint",
    "open-coverage"
)

if (!$command) {
    $command = "build-and-install"
}

if ($help) {
    Get-Help $MyInvocation.InvocationName
    exit
}

switch ($command) {
    {$script:availableCommands -contains $_} {
        & $command @Args
        break
    }
    default {
        Write-Output (
            """$command"" is not a valid command.",
            "To see available commands, type: ""$($MyInvocation.InvocationName) -help"""
        )
        break
    }
}
