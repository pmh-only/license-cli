#!/usr/bin/env pwsh
$ErrorActionPreference = 'Stop'
$install_path = "$Home\.license\bin"
$install_uri = 'https://raw.github.com/pmh-only/license-cli/master/build/windows_amd64.exe'

[Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12

if (!(Test-Path $install_path)) {
  New-Item $install_path -ItemType Directory | Out-Null
}

Invoke-WebRequest $install_uri -OutFile "$install_path\license.exe" -UseBasicParsing

$User = [EnvironmentVariableTarget]::User
$Path = [Environment]::GetEnvironmentVariable('Path', $User)

[Environment]::SetEnvironmentVariable('Path', "$Path;$install_path", $User)
$Env:Path += ";$install_path"

Write-Output "License CLI has been installed to $install_path."
Write-Output "Run 'license' to get started."