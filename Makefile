
VERSION := $(shell cat VERSION)
ARCHS := "linux/amd64 linux/arm linux/arm64 darwin/amd64"
GLIDE := $(shell command -v glide 2> /dev/null)
CARPET := $(shell command -v go-carpet 2> /dev/null)
MT := $(shell command -v multitail 2> /dev/null)
PWD := $(shell cd .. && pwd)

default: build

examples:
	@go build examples/filesbydate.go
	@go build examples/filesbysize.go

setup:
	@echo installing tools...
ifndef GLIDE
	@brew install glide
endif

mocks:
	@mockery -name Countable

install:
	@glide install

build:
	@go build

.PHONY: test build release setup install watch lint mocks coverage eject examples
