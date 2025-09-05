BINARY_NAME=pdf2png
GO=go
GOFLAGS=-v
LDFLAGS=-s -w

GREEN  := $(shell tput -T screen setaf 2)
YELLOW := $(shell tput -T screen setaf 3)
RED    := $(shell tput -T screen setaf 1)
CYAN   := $(shell tput -T screen setaf 6)
RESET  := $(shell tput -T screen sgr0)

UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Windows_NT)
    BINARY_NAME := $(BINARY_NAME).exe
endif

.PHONY: all build clean test run install deps fmt vet help test-e2e
.PHONY: build-all build-linux build-windows build-darwin
.DEFAULT_GOAL := help

all: build

build:
	@echo "$(YELLOW)Building $(BINARY_NAME)...$(RESET)"
	@$(GO) build $(GOFLAGS) -ldflags="$(LDFLAGS)" -o $(BINARY_NAME) .
	@echo "$(GREEN)Build complete.$(RESET)"

clean:
	@echo "$(YELLOW)Cleaning project...$(RESET)"
	@$(GO) clean
	@rm -f $(BINARY_NAME) $(BINARY_NAME)-*
	@rm -rf output_* example test_output
	@echo "$(GREEN)Clean complete.$(RESET)"

run: build
	@echo "$(YELLOW)Running example...$(RESET)"
	@./$(BINARY_NAME) -i example.pdf -o output_example -v

test:
	@echo "\n$(YELLOW)Running All Tests...$(RESET)"
	@echo "---------------------------------"
	@echo "$(CYAN)1. Running Unit Tests...$(RESET)"
	@$(GO) test -v ./...
	@echo "\n$(CYAN)2. Running End-to-End (E2E) Tests...$(RESET)"
	@make --no-print-directory test-e2e
	@echo "---------------------------------"
	@echo "$(GREEN)All tests passed successfully!$(RESET)"

test-e2e: build
	@echo "  - Test: Successful conversion"
	@./$(BINARY_NAME) -i testdata/sample.pdf -o test_output > /dev/null 2>&1 && [ -f test_output/page_001.png ] \
		&& echo "    $(GREEN)PASS$(RESET)" \
		|| (echo "    $(RED)FAIL$(RESET)"; exit 1)
	@rm -rf test_output

	@echo "  - Test: Missing input file argument"
	@./$(BINARY_NAME) 2>&1 | grep -q "Error: input PDF file is required" \
		&& echo "    $(GREEN)PASS$(RESET)" \
		|| (echo "    $(RED)FAIL$(RESET)"; exit 1)

	@echo "  - Test: Non-existent input file"
	@./$(BINARY_NAME) -i non_existent_file.pdf 2>&1 | grep -q "Error: input file 'non_existent_file.pdf' does not exist" \
		&& echo "    $(GREEN)PASS$(RESET)" \
		|| (echo "    $(RED)FAIL$(RESET)"; exit 1)

deps:
	@echo "$(YELLOW)Tidying Go modules...$(RESET)"
	@$(GO) mod tidy
	@echo "$(GREEN)Dependencies are up-to-date.$(RESET)"

fmt:
	@echo "$(YELLOW)Formatting code...$(RESET)"
	@$(GO) fmt ./...

vet:
	@echo "$(YELLOW)Vetting code...$(RESET)"
	@$(GO) vet ./...

install: deps
	@echo "$(YELLOW)Installing $(BINARY_NAME) to GOPATH/bin...$(RESET)"
	@$(GO) install $(GOFLAGS)
	@echo "$(GREEN)Installation complete.$(RESET)"

build-all: build-linux build-windows build-darwin

build-linux:
	@echo "$(YELLOW)Building for Linux (amd64)...$(RESET)"
	@GOOS=linux GOARCH=amd64 $(GO) build -ldflags="$(LDFLAGS)" -o $(BINARY_NAME)-linux-amd64 .

build-windows:
	@echo "$(YELLOW)Building for Windows (amd64)...$(RESET)"
	@GOOS=windows GOARCH=amd64 $(GO) build -ldflags="$(LDFLAGS)" -o $(BINARY_NAME)-windows-amd64.exe .

build-darwin:
	@echo "$(YELLOW)Building for macOS (amd64 & arm64)...$(RESET)"
	@GOOS=darwin GOARCH=amd64 $(GO) build -ldflags="$(LDFLAGS)" -o $(BINARY_NAME)-darwin-amd64 .
	@GOOS=darwin GOARCH=arm64 $(GO) build -ldflags="$(LDFLAGS)" -o $(BINARY_NAME)-darwin-arm64 .

help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "$(YELLOW)Available targets:$(RESET)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  $(CYAN)%-15s$(RESET) %s\n", $$1, $$2}'
