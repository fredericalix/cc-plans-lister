# Makefile for cc-plans-lister

# Variables
BINARY_NAME=cc-plans-lister
BUILD_DIR=bin
CMD_DIR=cmd/cc-plans-lister
MAIN_FILE=$(CMD_DIR)/main.go

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=$(GOCMD) fmt

# Build flags
LDFLAGS=-ldflags "-X main.version=$(VERSION)"
VERSION?=1.0.0

.PHONY: all build clean test coverage fmt vet deps help install

# Default target
all: fmt vet test build

# Build the application
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_FILE)

# Clean build artifacts
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)
	rm -f *.md *.txt *.csv *.pdf
	rm -f coverage.out

# Run tests
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

# Run tests with coverage
coverage:
	@echo "Running tests with coverage..."
	$(GOTEST) -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

# Format Go code
fmt:
	@echo "Formatting code..."
	$(GOFMT) ./...

# Run go vet
vet:
	@echo "Running go vet..."
	$(GOCMD) vet ./...

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	$(GOMOD) download
	$(GOMOD) tidy

# Install the binary to GOPATH/bin
install: build
	@echo "Installing $(BINARY_NAME)..."
	cp $(BUILD_DIR)/$(BINARY_NAME) $(GOPATH)/bin/

# Cross-compile for multiple platforms
build-all: clean
	@echo "Building for multiple platforms..."
	@mkdir -p $(BUILD_DIR)
	
	# Linux AMD64
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(MAIN_FILE)
	
	# Darwin AMD64 (Intel Mac)
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(MAIN_FILE)
	
	# Darwin ARM64 (Apple Silicon Mac)
	GOOS=darwin GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 $(MAIN_FILE)
	
	# Windows AMD64
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(MAIN_FILE)

# Run the application with default parameters
run: build
	@echo "Running $(BINARY_NAME)..."
	./$(BUILD_DIR)/$(BINARY_NAME)

# Generate sample outputs (requires CLEVER_API_TOKEN)
samples: build
	@echo "Generating sample outputs..."
	@if [ -z "$(CLEVER_API_TOKEN)" ]; then \
		echo "Error: CLEVER_API_TOKEN environment variable is required"; \
		exit 1; \
	fi
	./$(BUILD_DIR)/$(BINARY_NAME) --format=markdown --output=sample.md
	./$(BUILD_DIR)/$(BINARY_NAME) --format=txt --output=sample.txt
	./$(BUILD_DIR)/$(BINARY_NAME) --format=csv --output=sample.csv
	./$(BUILD_DIR)/$(BINARY_NAME) --format=pdf --output=sample.pdf
	@echo "Sample files generated: sample.{md,txt,csv,pdf}"

# Development workflow: format, vet, test, build
dev: fmt vet test build
	@echo "Development build completed successfully!"

# Help target
help:
	@echo "Available targets:"
	@echo "  all         - Format, vet, test, and build (default)"
	@echo "  build       - Build the application"
	@echo "  build-all   - Cross-compile for multiple platforms"
	@echo "  clean       - Clean build artifacts and output files"
	@echo "  test        - Run tests"
	@echo "  coverage    - Run tests with coverage report"
	@echo "  fmt         - Format Go code"
	@echo "  vet         - Run go vet"
	@echo "  deps        - Download and tidy dependencies"
	@echo "  install     - Install binary to GOPATH/bin"
	@echo "  run         - Build and run the application"
	@echo "  samples     - Generate sample output files (requires CLEVER_API_TOKEN)"
	@echo "  dev         - Development workflow (fmt, vet, test, build)"
	@echo "  help        - Show this help message"