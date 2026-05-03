GO ?= go
BINARY ?= 6502
BUILD_DIR ?= build
BIN_DIR ?= $(BUILD_DIR)/bin
COVERAGE_FILE ?= coverage.out
LOG_FOLDER ?= $(BUILD_DIR)/.logs
COVERAGE_FOLDER ?= $(BUILD_DIR)/.coverage
BINARY_PATH := $(BIN_DIR)/$(BINARY)
COVERAGE_PROFILE := $(COVERAGE_FOLDER)/$(COVERAGE_FILE)
GENERATED_FILES := computer/instruction_string.go
BUILD_SOURCES := $(shell find computer programs -name '*.go' ! -name '*_test.go') main.go go.mod go.sum
TEST_SOURCES := $(shell find . -name '*.go') go.mod go.sum

.DEFAULT_GOAL := help

.PHONY: help run test test-verbose cover-html fmt vet tidy clean

help:
	@printf "Available targets:\n"
	@printf "  build         Build the emulator binary\n"
	@printf "  run           Run the emulator\n"
	@printf "  test          Run the test suite\n"
	@printf "  test-verbose  Run the test suite with verbose output\n"
	@printf "  coverage      Generate a coverage profile\n"
	@printf "  cover-html    Open the coverage report in HTML format\n"
	@printf "  generate      Run go generate for the computer package\n"
	@printf "  fmt           Format all Go code\n"
	@printf "  vet           Run go vet\n"
	@printf "  tidy          Tidy module dependencies\n"
	@printf "  clean         Remove generated files\n"

build: $(BINARY_PATH)

$(BINARY_PATH): $(GENERATED_FILES) $(BUILD_SOURCES)
	mkdir -p $(BIN_DIR)
	$(GO) build -o $@ .

run: $(BINARY_PATH)
	mkdir -p $(LOG_FOLDER)
	LOG_FOLDER=$(LOG_FOLDER) ./$(BINARY_PATH)

test:
	mkdir -p $(LOG_FOLDER)
	LOG_FOLDER=$(LOG_FOLDER) $(GO) test ./...

test-verbose:
	mkdir -p $(LOG_FOLDER)
	LOG_FOLDER=$(LOG_FOLDER) $(GO) test -v ./...

$(COVERAGE_PROFILE): $(GENERATED_FILES) $(TEST_SOURCES)
	mkdir -p $(LOG_FOLDER)
	mkdir -p $(COVERAGE_FOLDER)
	LOG_FOLDER=$(LOG_FOLDER) $(GO) test ./... -coverprofile=$@

coverage: $(COVERAGE_PROFILE)
	$(GO) tool cover -func=$<

cover-html: coverage
	$(GO) tool cover -html=$(COVERAGE_PROFILE)

generate: $(GENERATED_FILES)

$(GENERATED_FILES): computer/instructions.go
	$(GO) generate -v ./computer

fmt:
	$(GO) fmt ./...

vet:
	$(GO) vet ./...

tidy:
	$(GO) mod tidy

clean:
	rm -f $(BINARY) $(BINARY_PATH) $(COVERAGE_FILE)
	rm -f $(COVERAGE_PROFILE)
	rm -f $(LOG_FOLDER)/*.log
