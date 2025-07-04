# Variables
BINARY_NAME=algorithm-visualiser
BUILD_DIR=bin
MAIN_PATH=./cmd

# Go related variables
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=gofmt

# Build flags
LDFLAGS=-ldflags="-s -w"

.PHONY: all build clean run install fmt vet tidy help

# Default target
all: clean fmt vet build

# Build the application
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

# Build and run
run: build
	@echo "Running $(BINARY_NAME)..."
	@./$(BUILD_DIR)/$(BINARY_NAME)

# Quick development run (no build optimization)
dev:
	@echo "Running $(BINARY_NAME) in development mode..."
	$(GOCMD) run $(MAIN_PATH)

# Install to GOPATH/bin
install:
	@echo "Installing $(BINARY_NAME)..."
	$(GOCMD) install $(MAIN_PATH)
	@echo "Installation complete"

# Clean build artifacts
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out coverage.html
	@echo "Clean complete"

# Format code
fmt:
	@echo "Formatting code..."
	$(GOFMT) -s -w .
	@echo "Format complete"

# Vet code
vet:
	@echo "Vetting code..."
	$(GOCMD) vet ./...
	@echo "Vet complete"

# Tidy dependencies
tidy:
	@echo "Tidying dependencies..."
	$(GOMOD) tidy
	@echo "Tidy complete"

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	$(GOMOD) download
	@echo "Dependencies downloaded"

# Development build (faster, includes debug info and race detection)
dev-build:
	@echo "Building $(BINARY_NAME) for development..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) -race -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)
	@echo "Development build complete: $(BUILD_DIR)/$(BINARY_NAME)"

# Cross-compilation targets
build-linux:
	@echo "Building for Linux..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(MAIN_PATH)

build-windows:
	@echo "Building for Windows..."
	@mkdir -p $(BUILD_DIR)
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(MAIN_PATH)

build-darwin:
	@echo "Building for macOS (Intel)..."
	@mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(MAIN_PATH)

build-darwin-arm:
	@echo "Building for macOS (Apple Silicon)..."
	@mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 $(MAIN_PATH)

build-all: build-linux build-windows build-darwin build-darwin-arm

# Create release archives
release: clean build-all
	@echo "Creating release archives..."
	@mkdir -p $(BUILD_DIR)/release
	@tar -czf $(BUILD_DIR)/release/$(BINARY_NAME)-linux-amd64.tar.gz -C $(BUILD_DIR) $(BINARY_NAME)-linux-amd64 -C .. README.md
	@tar -czf $(BUILD_DIR)/release/$(BINARY_NAME)-darwin-amd64.tar.gz -C $(BUILD_DIR) $(BINARY_NAME)-darwin-amd64 -C .. README.md
	@tar -czf $(BUILD_DIR)/release/$(BINARY_NAME)-darwin-arm64.tar.gz -C $(BUILD_DIR) $(BINARY_NAME)-darwin-arm64 -C .. README.md
	@cd $(BUILD_DIR) && zip -j release/$(BINARY_NAME)-windows-amd64.zip $(BINARY_NAME)-windows-amd64.exe ../README.md
	@echo "Release archives created in $(BUILD_DIR)/release/"

# Show supported algorithms
algorithms:
	@echo "Supported sorting algorithms:"
	@echo "  1. Bubble Sort    - Basic comparison-based sorting"
	@echo "  2. Selection Sort - Find minimum and swap"

# Quality checks (combines formatting, and vetting)
check: fmt vet
	@echo "All quality checks passed!"

# Help
help:
	@echo "Available targets:"
	@echo "  all           - Clean, format, vet, and build (default)"
	@echo "  build         - Build the application (optimized)"
	@echo "  dev-build     - Build with race detection for development"
	@echo "  run           - Build and run the application"
	@echo "  dev           - Run directly without building (fastest for development)"
	@echo "  install       - Install to GOPATH/bin"
	@echo "  clean         - Remove build artifacts"
	@echo "  fmt           - Format code"
	@echo "  vet           - Vet code"
	@echo "  tidy          - Tidy dependencies"
	@echo "  deps          - Download dependencies"
	@echo "  build-linux   - Cross-compile for Linux"
	@echo "  build-windows - Cross-compile for Windows"
	@echo "  build-darwin  - Cross-compile for macOS (Intel)"
	@echo "  build-darwin-arm - Cross-compile for macOS (Apple Silicon)"
	@echo "  build-all     - Cross-compile for all platforms"
	@echo "  release       - Create release archives"
	@echo "  algorithms    - Show supported algorithms and controls"
	@echo "  check         - Run fmt and vet"
	@echo "  help          - Show this help message"