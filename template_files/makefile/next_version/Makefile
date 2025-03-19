# cdcloud.dev Makefile for GO.
# author: cd-stephen
MAKE_FILE_VER := 1.0.5.beta

# INFO
# @         | only show the output of the command, not the command itself
# -         | Ignores errors from the command. ie @-<command> means silent no stderr out on command
# ./...     | instructs go tool to apply the command to all files in local and subdirectories
# printf -- | instructs printf to interpret -- as string and not options

## Module info
BIN_DIR := bin
MODULE_NAME := <project/module name>
URL_PATH := <github.com/... url>

## Load variables from a .env file.
-include .env

## Path variable
GOBASE ?= $(CURDIR)

## Build variables
VERSION := $(shell git describe --tags 2>/dev/null || echo "v0.0.1")
BUILD := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
PROJECTNAME := $(shell basename "$(PWD)")

## Embed version info in binaries #CHANGE MADE HERE - FROM: none TO: Added embedded version flags
LDFLAGS := -s -w -X main.Version=$(VERSION) -X main.Commit=$(BUILD)

## Color variables (new) #CHANGE MADE HERE - FROM: none TO: Added colorized output variables
RED := $(shell tput setaf 1)
GREEN := $(shell tput setaf 2)
YELLOW := $(shell tput setaf 3)
RESET := $(shell tput sgr0)

## Default (no flags)
.DEFAULT_GOAL := help

# .PHONY as targets do not represent files.
.PHONY: help fmt vet lint run build build-prod build-proto test test-with-cover generate-mocks clean deps tidy asm

# Print header information and available targets
help:
	@printf "\n"
	@printf "##########################################################\n"
	@printf "#          📘 cdcloud.io GO build tool 📘                #\n"
	@printf "#               version: $(MAKE_FILE_VER)                      #\n"
	@printf "##########################################################\n"
	@printf "\n"
	@printf "📄 Makefile Variables 📄 \n"
	@printf -- "------------------------\n"
	@printf "VERSION: %s\n" "$(VERSION)"
	@printf "BUILD: %s\n" "$(BUILD)"
	@printf "BIN_DIR: %s\n" "$(BIN_DIR)"
	@printf "MODULE_NAME: %s\n" "$(MODULE_NAME)"
	@printf "URL_PATH: %s\n" "$(URL_PATH)"
	@printf "GOBASE: %s\n" "$(GOBASE)"
	@printf "\n"
	@printf "🎯 Available Targets 🎯 \n"
	@printf -- "-----------------------\n"
	@printf "help              - Show this help information\n"
	@printf "fmt               - Format the code using gofmt\n"
	@printf "vet               - Run go vet to check for issues\n"
	@printf "lint              - Run golangci-lint to analyze the code\n"
	@printf "build             - Build the project\n"
	@printf "build-prod        - Build the project for production with stripped symbols\n"
	@printf "build-proto       - Build protobuf definitions\n"
	@printf "run               - Run the application\n"
	@printf "test              - Run the tests\n"
	@printf "test-with-cover   - Run tests with coverage\n"
	@printf "generate-mocks    - Generate mocks using mockery\n"
	@printf "clean             - Clean the project and build artifacts\n"
	@printf "deps              - Get dependencies\n"
	@printf "tidy              - Tidy and download Go modules\n"
	@printf "asm               - Generate assembly code\n"
	@printf "\n"

# Ensure BIN_DIR exists
$(BIN_DIR):
	@mkdir -p $@

vars:
	@printf "\n"
	@printf "##########################################################\n"
	@printf "#          📘  cdcloud.io Makefile VAR info 📘           #\n"
	@printf "##########################################################\n"
	@printf "\n"
	@printf "📄 Makefile variables 📄 \n"
	@printf -- "-----------------------\n"
	@printf "VERSION: %s\n" "$(VERSION)"
	@printf "BUILD: %s\n" "$(BUILD)"
	@printf "BIN_DIR: %s\n" "$(BIN_DIR)"
	@printf "MODULE_NAME: %s\n" "$(MODULE_NAME)"
	@printf "URL_PATH: %s\n" "$(URL_PATH)"
	@printf "GOBASE: %s\n" "$(GOBASE)"
	@printf "\n"

# Formats the code to the official GO style guideline
fmt:
	@echo "$(YELLOW)Formatting code...$(RESET)" #CHANGE MADE HERE - FROM: silent TO: colorized status message
	@go fmt ./...

# Analyzes GO code for potential errors or suspicious constructs, but does not modify the code.
vet:
	@echo "$(YELLOW)Running go vet...$(RESET)" #CHANGE MADE HERE - FROM: silent TO: colorized status message
	@go vet ./...

# Lint analyzes source code to identify potential issues. (untested)
lint: check-tools
	@echo "$(YELLOW)Running golangci-lint...$(RESET)" #CHANGE MADE HERE - FROM: silent TO: colorized status message
	@golangci-lint run --enable-all

# Check tools (new target to verify required tools exist)
check-tools:
	@command -v golangci-lint >/dev/null || (echo "$(RED)golangci-lint is not installed!$(RESET)"; exit 1)
	@command -v mockery >/dev/null || (echo "$(RED)mockery is not installed!$(RESET)"; exit 1)

# Run the application
run:
	@echo "$(GREEN)Running application...$(RESET)" #CHANGE MADE HERE - FROM: silent TO: colorized status message
	@go run cmd/$(MODULE_NAME)/*.go

# Run the qagent
run-qagent:
	@echo "$(GREEN)Running qagent...$(RESET)" #CHANGE MADE HERE - FROM: silent TO: colorized status message
	@go run cmd/qagent-dreamstream/*.go

# Generic build target (new), allows specifying BUILD_TARGET on command line
BUILD_TARGET ?= $(MODULE_NAME)

# Build the application
build: deps | $(BIN_DIR)
	@echo "$(GREEN)Building $(BUILD_TARGET)...$(RESET)" #CHANGE MADE HERE - FROM: simple message TO: colorized and dynamic target message
	@go build -ldflags="$(LDFLAGS)" -o $(BIN_DIR)/$(BUILD_TARGET) ./cmd/$(BUILD_TARGET)

# Build the application for production delivery (strip symbols)
build-prod: deps | $(BIN_DIR)
	@echo "$(GREEN)Building production binary $(BUILD_TARGET)...$(RESET)" #CHANGE MADE HERE - FROM: simple message TO: colorized and dynamic target message
	@go build -mod=vendor -ldflags="$(LDFLAGS)" -o $(BIN_DIR)/$(BUILD_TARGET)/$(BUILD_TARGET) ./cmd/$(BUILD_TARGET) || (echo "$(RED)Build failed with exit code $$?$(RESET)"; exit 1)

# Build proto-buf code (untested)
build-proto:
	@echo "$(YELLOW)Building protobuf files...$(RESET)" #CHANGE MADE HERE - FROM: silent TO: colorized status message
	@-protoc --proto_path=protobuf protobuf/*.proto --go_out=internal --go-grpc_out=internal

# Run tests (untested)
test:
	@echo "$(YELLOW)Running tests...$(RESET)" #CHANGE MADE HERE - FROM: silent TO: colorized status message
	@go test -v $(shell go list ./... | grep -v /test/)

# Run tests with coverage (untested)
test-with-cover:
	@echo "$(YELLOW)Running tests with coverage...$(RESET)" #CHANGE MADE HERE - FROM: silent TO: colorized status message
	@go test -v -coverprofile=cover.out $(shell go list ./... | grep -v /test/)
	@go tool cover -html=cover.out -o cover.html
	@echo "$(GREEN)Coverage report generated: cover.html$(RESET)"

# Generate mocks (untested)
generate-mocks: check-tools
	@echo "$(YELLOW)Generating mocks...$(RESET)" #CHANGE MADE HERE - FROM: silent TO: colorized status message
	@-mockery --all --with-expecter --keeptree

# Clean the build artifacts
clean:
	@echo "$(RED)Cleaning build artifacts...$(RESET)" #CHANGE MADE HERE - FROM: silent TO: colorized status message
	@go clean
	@-rm -rf $(BIN_DIR)
	@-rm -rf vendor
	@-rm -f cover.out cover.html

# Clean the mod cache folder - If you really want to nuke everything
clean-all:
	@echo "$(RED)Performing deep clean (mod cache, etc)...$(RESET)" #CHANGE MADE HERE - FROM: silent TO: colorized status message
	@go clean -cache -modcache -i -r
	@-rm -rf $(BIN_DIR)
	@-rm -rf vendor
	@-rm -f cover.out cover.html

# Dependency management
deps: tidy
	@echo "$(YELLOW)Downloading dependencies...$(RESET)" #CHANGE MADE HERE - FROM: silent TO: colorized status message
	@go mod download
	@go mod vendor # new: ensure vendor is populated for -mod=vendor builds

# Tidy and update modules
tidy:
	@echo "$(YELLOW)Tidying modules...$(RESET)" #CHANGE MADE HERE - FROM: silent TO: colorized status message
	@go mod tidy

# Generate assembly code (untested)
asm:
	@echo "$(YELLOW)Generating assembly code...$(RESET)" #CHANGE MADE HERE - FROM: silent TO: colorized status message
	@go tool compile -S cmd/$(MODULE_NAME)/main.go > $(MODULE_NAME).asm