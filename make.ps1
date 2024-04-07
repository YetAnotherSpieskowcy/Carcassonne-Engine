<#
.Synopsis
Makefile script in PowerShell that contains commands useful during development of Carcassonne-Engine.

.Description
Available commands:
   build             Build all Go source files.
   test              Run the test suite.
   open-coverage     Show coverage in the browser after running the test suite.

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
        $script:availableCommands = @("build", "test", "open-coverage")
        return $script:availableCommands | Where-Object { $_ -like "$wordToComplete*" }
    })]
    [String]
    $command,
    [switch]
    $help = $false
)

function build() {
    Write-Output "Building the project..."
    & go build "./..."
    Exit $LASTEXITCODE
}

function test() {
    Write-Output "Running the test suite..."
    & go test -race "-coverprofile=coverage.txt" "./..."
    Exit $LASTEXITCODE
}

function open-coverage() {
    & go tool cover "-html=coverage.txt"
    Write-Output "Coverage opened in the default browser."
}

$script:availableCommands = @("build", "test", "open-coverage")

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
