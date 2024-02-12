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

# For macOS, we use a single binary called 'macos'
if [ "$OS" = "macos" ]; then
  BINARY="./tw/macos"
else
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
fi

echo $BINARY

# Infer pages from .html files in the pages directory
PAGES=$(ls ../pages/templates/*.html | xargs -n 1 basename | cut -d. -f1)

# Run the binary for each page
for PAGE in $PAGES; do
  (
    # Detect which partials are being used in this page
    PARTIALS=$(grep -o -E '{{template "[^"]+' ../pages/templates/${PAGE}.html | cut -d'"' -f2 | xargs -I{} echo \"../pages/templates/partials/{}.html\")

    # Generate an array of partials and join them with commas
    PARTIALS_ARRAY=$(echo $PARTIALS | tr ' ' ',')

    # Always include the "header" partial and any other partials that are always used
    PARTIALS_ARRAY=\"../pages/templates/partials/header.html\",\"../pages/templates/partials/global.html\",$PARTIALS_ARRAY

    # Generate Tailwind config for this page
    echo "module.exports = {
      content: [\"../pages/templates/${PAGE}.html\", \"../pages/templates/layouts/*.html\", $PARTIALS_ARRAY],
      theme: {
        extend: {},
      },
      daisyui: {
        themes: [\"night\"],
      },
      plugins: [require('daisyui'), require('@tailwindcss/typography')],
    }" > tailwind.config.${PAGE}.js

    # Run the binary with the generated config
    $BINARY build -i ./base.css -c tailwind.config.${PAGE}.js -o ../public/css/styles.${PAGE}.css --minify
  ) &
done

# Wait for all background processes to finish
wait