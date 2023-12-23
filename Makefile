run:
	@templ generate
	@go run cmd/main.go

.PHONY: build-css
build-css: tailwindcss
	./tailwindcss -i tailwind.css -o assets/styles/tailwind.css --minify

# see for details -> https://tailwindcss.com/blog/standalone-cli
TAILWIND_RELEASE = tailwindcss-macos-x64
tailwindcss:
	curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/$(TAILWIND_RELEASE)
	chmod +x $(TAILWIND_RELEASE)
	mv $(TAILWIND_RELEASE) tailwindcss

.PHONY: build-css
watch-css: tailwindcss
	./tailwindcss -i tailwind.css -o assets/styles/tailwind.css --watch
