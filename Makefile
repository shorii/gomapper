.DEFAULT_GOAL := help

## Install dependencies
.PHONY: deps
deps:
	go get -v -d

## Setup
.PHONY: deps
devel-deps: deps
	GO111MODULE=off go get \
	golang.org/x/lint \
	github.com/Songmu/make2help/cmd/make2help

## Run tests
.PHONY: test
test: deps
	./_tools/go_test.sh

## Show help
.PHONY: help
help:
	@make2help $(MAKEFILE_LIST)
