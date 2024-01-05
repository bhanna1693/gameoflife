run:
	air

air-local:
	make build-tailwind
	make build-templ
	make build-app

build-app:
	go build -o ./tmp/main ./cmd

build-templ:
	templ generate

.PHONY: build-tailwind
build-tailwind: tailwindcss
	./tailwindcss -i tailwind.css -o web/static/styles/tailwind.css --minify

# see for details -> https://tailwindcss.com/blog/standalone-cli
TAILWIND_RELEASE = tailwindcss-macos-x64
fetch-tailwindcss:
	curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/$(TAILWIND_RELEASE)
	chmod +x $(TAILWIND_RELEASE)
	mv $(TAILWIND_RELEASE) tailwindcss
