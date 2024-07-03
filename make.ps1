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
        $script:availableCommands = @("build", "test", "open-coverage", "lint")
        return $script:availableCommands | Where-Object { $_ -like "$wordToComplete*" }
    })]
    [String]
    $command,
    [switch]
    $help = $false
)

function build() {
    if (!(Test-Path .venv)) {
        & py -3.12 -m venv .venv
    }
    Write-Output "Building the project..."
    & go build "./pkg/..."
    if ($LASTEXITCODE) {
        Exit $LASTEXITCODE
    }
    & .venv\Scripts\python.exe -m pip install .
    Exit $LASTEXITCODE
}

function test() {
    Write-Output "Running the test suite..."
    & go test -race "-coverprofile=coverage.txt" "./pkg/..."
    Exit $LASTEXITCODE
}

function open-coverage() {
    & go tool cover "-html=coverage.txt"
    Write-Output "Coverage opened in the default browser."
}

function lint() {
    Write-Output "Running the linter..."
    & docker run -e "VALIDATE_ALL_CODEBASE=true" -e "DEFAULT_BRANCH=origin/main" -e "VALIDATE_GO=false" -e "LOG_LEVEL=NOTICE" -e "RUN_LOCAL=true" -v ".:/tmp/lint" --rm "ghcr.io/super-linter/super-linter:v6.3.0"
}

$script:availableCommands = @("build", "test", "open-coverage", "lint")

if (!$command) {
    $command = "build"
}

if ($help) {
    Get-Help $MyInvocation.InvocationName
    exit
}

switch ($command) {
    {$script:availableCommands -contains $_} {
        & $command
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
