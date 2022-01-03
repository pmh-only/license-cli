#!/bin/sh
if [ "$(id -un)" != "root" ]; then
  echo "This script must be run as root (sudo)"
  exit 1
fi

if command -v license > /dev/null; then
  rm "$(which license)"
fi

case $(uname -sm) in
  "Darwin x86_64") target="darwin_amd64" ;;
  "Linux x86_64") target="linux_amd64" ;;
  "Linux i686") target="linux_386" ;;
esac

install_uri="https://github.com/pmh-only/license-cli/releases/download/v0.1/license-cli_${target}"
install_path="/usr/local/bin/license-cli"

curl --fail --location --progress-bar --output "${install_path}" "${install_uri}"
chmod a+x "${install_path}"

echo "License CLI has been installed to ${install_path}."
echo "Run 'license' to get started."