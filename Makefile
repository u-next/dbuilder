NAME := dbuilder
VERSION := 1.0.0

.DEFAULT_GOAL := help

.PHONY: help
help: ## show help see: https://postd.cc/auto-documented-makefile/
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: test
test: ## run tests
	go test ./...
	golangci-lint run
