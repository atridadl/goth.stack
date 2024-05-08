#!/bin/bash

# Parse command-line arguments
while [[ $# -gt 0 ]]; do
  case $1 in
    -e|--extensions)
      EXTENSIONS="$2"
      shift; shift
      ;;
    -d|--directory)
      DIRECTORY="$2"
      shift; shift
      ;;
    -o|--output-dir)
      OUTPUT_DIR="$2"
      shift; shift
      ;;
    *)
      echo "Unknown argument: $1"
      exit 1
      ;;
  esac
done

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

echo "Extensions: $EXTENSIONS"
echo "Directory: $DIRECTORY"
echo "Output Directory: $OUTPUT_DIR"

# Set default values if not provided
OUTPUT_DIR="${OUTPUT_DIR:-../../public/css}"
DIRECTORY="${DIRECTORY:-.}"

if [[ -z "$EXTENSIONS" ]]; then
    echo "No extensions provided."
    exit 1
fi

# Initialize an array for name conditions
name_conditions=()

# Assuming $EXTENSIONS is a comma-separated list of extensions
IFS=',' read -ra ADDR <<< "$EXTENSIONS"
for ext in "${ADDR[@]}"; do
    name_conditions+=(-name "*.$ext")
done

# Use find with the array of conditions
INCLUDE_FILES=$(find "$DIRECTORY" -type f \( "${name_conditions[@]}" \))

echo "Files found: $INCLUDE_FILES"

# Optionally, remove leading './' if necessary
INCLUDE_FILES=$(echo "$INCLUDE_FILES" | sed 's|^\./||')

INCLUDE_FILES_ARRAY=$(echo "$INCLUDE_FILES" | awk '{printf "\"%s\",", $0}' | sed 's/,$//')

# Generate Tailwind config
echo "module.exports = {
  content: [$INCLUDE_FILES_ARRAY],
  theme: {
    extend: {},
  },
  daisyui: {
    themes: [\"night\"],
  },
  plugins: [require('daisyui'), require('@tailwindcss/typography')],
}" > tailwind.config.js

# Run the binary with the generated config
$BINARY build -i ./base.css -c tailwind.config.js -o "${OUTPUT_DIR}/styles.css" --minify

# Wait for all background processes to finish
wait