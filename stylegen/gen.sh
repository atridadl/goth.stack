#!/bin/sh

OS=$(uname -s)
ARCH=$(uname -m)

# Normalize OS and ARCH identifiers
case $OS in
  "Darwin")
    OS="macos"
    ;;
  "Linux")
    OS="linux"
    ;;
  "CYGWIN"*|"MINGW"*|"MSYS"*)
    OS="windows"
    ;;
  *)
    echo "Unknown operating system: $OS"
    exit 1
    ;;
esac

case $ARCH in
  "x86_64")
    ARCH="x64"
    ;;
  "arm64")
    ARCH="arm64"
    ;;
  *)
    echo "Unsupported architecture: $ARCH"
    exit 1
    ;;
esac

# Construct the binary file name
BINARY="./tw/${OS}-${ARCH}"
if [ "$OS" = "windows" ]; then
  BINARY="${BINARY}.exe"
else
  # Set execute permissions on the binary
  chmod +x $BINARY
fi

# Run the binary
$BINARY build -i ./base.css -o ../public/css/styles.css --minify