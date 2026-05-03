GO ?= go
BINARY ?= 6502
COVERAGE_FILE ?= coverage.out
LOG_FOLDER ?= .logs
COVERAGE_FOLDER ?= .coverage

.DEFAULT_GOAL := help

.PHONY: help build run test test-verbose coverage cover-html fmt vet tidy clean

help:
	@printf "Available targets:\n"
	@printf "  build         Build the emulator binary\n"
	@printf "  run           Run the emulator\n"
	@printf "  test          Run the test suite\n"
	@printf "  test-verbose  Run the test suite with verbose output\n"
	@printf "  coverage      Generate a coverage profile\n"
	@printf "  cover-html    Open the coverage report in HTML format\n"
	@printf "  fmt           Format all Go code\n"
	@printf "  vet           Run go vet\n"
	@printf "  tidy          Tidy module dependencies\n"
	@printf "  clean         Remove generated files\n"

build:
	$(GO) build -o $(BINARY) .

run:
	LOG_FOLDER=$(LOG_FOLDER) $(GO) run .

test:
	LOG_FOLDER=$(LOG_FOLDER) $(GO) test ./...

test-verbose:
	LOG_FOLDER=$(LOG_FOLDER) $(GO) test -v ./...

coverage:
	mkdir -p $(COVERAGE_FOLDER)
	LOG_FOLDER=$(LOG_FOLDER) $(GO) test ./... -coverprofile=$(COVERAGE_FOLDER)/$(COVERAGE_FILE)
	$(GO) tool cover -func=$(COVERAGE_FOLDER)/$(COVERAGE_FILE)

cover-html: coverage
	$(GO) tool cover -html=$(COVERAGE_FOLDER)/$(COVERAGE_FILE)

fmt:
	$(GO) fmt ./...

vet:
	$(GO) vet ./...

tidy:
	$(GO) mod tidy

clean:
	rm -f $(BINARY) $(COVERAGE_FILE)
	rm -f $(COVERAGE_FOLDER)/$(COVERAGE_FILE)
	rm -rf .logs
	rm -rf .coverage
