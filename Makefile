APP_NAME := $(shell grep -lR "func main()" *.go | awk -F/ '{print $$NF}' | sed 's/\.go//')
PACKAGES := $(shell go list ./...)
name := $(shell basename ${PWD})

all: help

.PHONY: help
help: Makefile
	@echo
	@echo " Application Name: $(APP_NAME)"
	@echo
	@echo " Choose a make command to run"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

## vet: vet code
.PHONY: vet
vet:
	go vet $(PACKAGES)

## test: run unit tests
.PHONY: test
test:
	go test -race -cover $(PACKAGES)

## css: build tailwindcss
.PHONY: css
css:
	npx tailwindcss -i assets/css/input.css -o assets/css/output.css --minify

## css-watch: watch build tailwindcss
.PHONY: css-watch
css-watch:
	npx tailwindcss -i assets/css/input.css -o assets/css/output.css --watch

## templ-generate: generate new template
.PHONY: templ-generate
templ-generate: 
	templ generate

## templ-watch: watch templ files and format them
.PHONY: templ-watch
templ-watch: 
	templ generate --watch

## staticcheck: run staticcheck
.PHONY: staticcheck
staticcheck:
	staticcheck ./...
