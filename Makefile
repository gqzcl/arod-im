GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
COMMON_SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
ifeq ($(origin ROOT_DIR),undefined)
ROOT_DIR := $(abspath $(shell cd $(COMMON_SELF_DIR) && pwd -P))
endif

# -------------------------------------------------------------------------
# Includes

include scripts/make-rules/tools.mk
include scripts/make-rules/init.mk
include scripts/make-rules/gen_config.mk
include scripts/make-rules/gen_api.mk
include scripts/make-rules/gen_wire.mk
include scripts/make-rules/run.mk
include scripts/make-rules/copyright.mk


# -------------------------------------------------------------------------

## all: generate all
.PHONY: all
all: init api config generate

## init: Initialize project dependencies
.PHONY: init
init:
	@$(MAKE) init.install

## config: Generate Configuration Files and You can use such as [ make config.job ] to generate the configuration file of a module separately
.PHONY: config
config:
	@$(MAKE) config.all

## api: Generate proto file of API
.PHONY: api
api:
	@$(MAKE) api.all

## generate: generate
.PHONY: generate
generate:
	go mod tidy
	go get github.com/google/wire/cmd/wire@latest
	go generate ./...

## wire: Generate dependency injection file and You can use such as [ make wire.job ] to generate the injection file of a module separately
.PHONY: wire
wire:
	@$(MAKE) wire.all

## run: Start the service using such as the [ make run.job ] command

## build: Build executable
.PHONY: build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...


## tools: install dependent tools,such as < make tools.install.swagger >
.PHONY: tools
tools:
	@$(MAKE) tools.install

## copyright-verify: Verify whether all files have copyright added.
.PHONY: copyright-verify
copyright-verify:
	@$(MAKE) copyright.verify

## copyright-add: Ensures source code files have copyright license headers.
.PHONY: copyright-add
copyright-add:
	@$(MAKE) copyright.add



## help: show help
.PHONY: help
help: Makefile
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@echo "$$USAGE_OPTIONS"

# 如果没有指定目标，则构建help目标
.DEFAULT_GOAL := help
