#!/bin/sh

TAILWIND_RELEASE="tailwindcss-macos-x64"

curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/$TAILWIND_RELEASE
chmod +x $TAILWIND_RELEASE
mv $TAILWIND_RELEASE tailwindcss
./tailwindcss -i ./../tailwind.css -o ./../web/static/styles/tailwind.css --minify
templ generate
go build -o ./tmp/main ./cmd