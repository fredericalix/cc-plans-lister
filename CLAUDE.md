# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

CC Plans Lister is a comprehensive Go CLI application that fetches data from the Clever Cloud API to generate documentation of available addon providers and application instance types. The tool supports multiple output formats including Markdown, plain text, CSV, and PDF.

## Commands

### Building and Running
```bash
# Build the application
make build
# or
go build -o bin/cc-plans-lister cmd/cc-plans-lister/main.go

# Run the application
./bin/cc-plans-lister --help
./bin/cc-plans-lister --format=markdown --output=services.md

# Cross-platform builds
make build-all
```

### Development Commands
```bash
# Development workflow (format, vet, test, build)
make dev

# Format code
make fmt
go fmt ./...

# Run tests
make test
go test ./...

# Run tests with coverage
make coverage

# Check for issues
make vet
go vet ./...

# Install dependencies
make deps
go mod tidy
```

### Usage Examples
```bash
# Environment setup
export CLEVER_API_TOKEN="your_token_here"

# Generate different output formats
./bin/cc-plans-lister --format=markdown --output=services.md
./bin/cc-plans-lister --format=txt --output=services.txt
./bin/cc-plans-lister --format=csv --output=services.csv
./bin/cc-plans-lister --format=pdf --output=services.pdf

# Output to stdout
./bin/cc-plans-lister --format=markdown
```

## Architecture

The application follows a modular architecture with clean separation of concerns:

### Directory Structure
- `cmd/cc-plans-lister/` - Main application entry point with CLI handling
- `internal/` - Private application packages
  - `api/` - Clever Cloud API client wrapper
  - `config/` - Configuration management and validation
  - `formatters/` - Output format implementations (markdown, txt, csv, pdf)
- `pkg/clevercloud/` - Public types and data structures
- `test/fixtures/` - Test data and utilities

### Core Components

**CLI Interface (`cmd/cc-plans-lister/main.go`)**
- Uses Cobra framework for command-line parsing
- Supports `--format` and `--output` flags
- Handles authentication token validation
- Manages error reporting and help messages

**API Client (`internal/api/client.go`)**
- Wraps the official `go.clever-cloud.dev/client` library
- Handles OAuth authentication via environment variables
- Provides methods: `GetAddonProviders()` and `GetProductInstances()`
- Includes data sorting for consistent output

**Configuration (`internal/config/config.go`)**
- Loads and validates `CLEVER_API_TOKEN` environment variable
- Validates supported output formats
- Provides configuration defaults

**Formatters (`internal/formatters/`)**
- Interface-based design with `Formatter` interface
- Four implementations: `MarkdownFormatter`, `TextFormatter`, `CSVFormatter`, `PDFFormatter`
- Each formatter generates comprehensive reports with:
  - Summary tables for addons and applications
  - Detailed plans with pricing and specifications
  - Grouped sections by provider/application type

### Data Flow
1. CLI validates arguments and loads configuration
2. API client authenticates and fetches data from two endpoints:
   - `/v2/products/addonproviders` - Addon providers and plans
   - `/v2/products/instances` - Application types and flavors
3. Data is sorted for consistent output
4. Selected formatter processes data and writes to specified output
5. Success/error reporting to user

## Key Dependencies

- **CLI Framework**: `github.com/spf13/cobra` - Command-line interface
- **PDF Generation**: `github.com/signintech/gopdf` - PDF document creation
- **Testing**: `github.com/stretchr/testify` - Test assertions and utilities
- **API Client**: `go.clever-cloud.dev/client` - Official Clever Cloud Go client
- **Standard Library**: Extensive use of `encoding/csv`, `text/tabwriter`, `sort`, etc.

## Authentication

The application requires authentication via the `CLEVER_API_TOKEN` environment variable:

```bash
export CLEVER_API_TOKEN="your_api_token_here"
```

The token is validated at startup and used for OAuth authentication with the Clever Cloud API.

## Testing

The project includes comprehensive unit tests:
- Configuration validation tests
- Formatter output tests with fixture data  
- API client initialization tests
- Utility function tests (text truncation, etc.)

Run tests with `make test` or `go test ./...`

## Common Issues

**Font Issues with PDF Generation**
- PDF formatter includes fallback logic for font loading
- Test environment uses minimal PDF output when fonts unavailable
- Production usage may require TTF font files in `./fonts/` directory

**API Rate Limiting**
- The Clever Cloud API may have rate limits
- Consider implementing retry logic for production usage

**Large Data Sets**
- Current implementation loads all data into memory
- For very large datasets, consider streaming or pagination