MAKEFLAGS := --no-print-directory --silent

default: help

help:
	@echo "Please use 'make <target>' where <target> is one of"
	awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z\._-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

terror: test.errors
test.errors: ## Run tests, alias: t.all
	go test ./... -run TestResponseError

tall: test.all
test.all: ## Run tests, alias: t.all
	go test ./...
