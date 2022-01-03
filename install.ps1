#!/usr/bin/env pwsh
$ErrorActionPreference = 'Stop'
$install_path = "$Home\.license\bin"

[Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12

if (!(Test-Path $install_path)) {
  New-Item $install_path -ItemType Directory | Out-Null
}

if ((gwmi win32_operatingsystem | select osarchitecture).osarchitecture -eq "64-bit") {
  Invoke-WebRequest 'https://github.com/pmh-only/license-cli/releases/download/v0.1/windows_amd64.exe' -OutFile "$install_path\license.exe" -UseBasicParsing
} else {
  Invoke-WebRequest 'https://github.com/pmh-only/license-cli/releases/download/v0.1/windows_386.exe' -OutFile "$install_path\license.exe" -UseBasicParsing
}

$User = [EnvironmentVariableTarget]::User
$Path = [Environment]::GetEnvironmentVariable('Path', $User)

[Environment]::SetEnvironmentVariable('Path', "$Path;$install_path", $User)
$Env:Path += ";$install_path"

Write-Output "License CLI has been installed to $install_path."
Write-Output "Run 'license' to get started."