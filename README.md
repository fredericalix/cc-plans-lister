# CC Plans Lister

A command-line tool that fetches and documents all available addon providers and application instance types from the Clever Cloud API. Generate comprehensive reports in multiple formats including Markdown, plain text, CSV, and PDF.

## Features

- **Multi-format output**: Support for Markdown, plain text, CSV, and PDF formats
- **Comprehensive data**: Lists all addon providers with their plans and application types with their flavors
- **Structured information**: Organized tables with pricing, specifications, and availability
- **CLI interface**: Easy-to-use command-line interface with flexible options
- **API integration**: Direct integration with Clever Cloud's official API

## Installation

### Prerequisites

- Go 1.23 or higher
- Valid Clever Cloud API token

### Build from source

```bash
git clone <repository-url>
cd cc-plans-lister
make deps    # Download and install dependencies
make build   # Build the application
```

### Quick start

```bash
# Build, test, and prepare for development
make all     # Runs: format, vet, test, build

# Or use the development workflow
make dev     # Same as 'make all' with success message
```

## Configuration

### Authentication

The tool requires a Clever Cloud API token. Set the environment variable:

```bash
export CLEVER_API_TOKEN="your_api_token_here"
```

You can obtain an API token from your Clever Cloud console.

## Usage

### Basic usage

```bash
# Generate markdown output to stdout
./bin/cc-plans-lister

# Specify output format
./bin/cc-plans-lister --format=txt

# Save to file
./bin/cc-plans-lister --format=pdf --output=clever-cloud-services.pdf
```

### Command-line options

```
Usage:
  cc-plans-lister [flags]
  cc-plans-lister [command]

Available Commands:
  help        Help about any command
  version     Print the version number

Flags:
  -f, --format string   Output format (markdown, txt, csv, pdf) (default "markdown")
  -h, --help           help for cc-plans-lister
  -o, --output string   Output file (default: stdout)
```

### Output formats

#### Markdown (default)
```bash
./bin/cc-plans-lister --format=markdown --output=services.md
```
Generates a well-formatted markdown document with tables and sections.

#### Plain Text
```bash
./bin/cc-plans-lister --format=txt --output=services.txt
```
Creates a plain text report with tabular data suitable for terminal viewing.

#### CSV
```bash
./bin/cc-plans-lister --format=csv --output=services.csv
```
Exports structured data in CSV format for spreadsheet applications.

#### PDF
```bash
./bin/cc-plans-lister --format=pdf --output=services.pdf
```
Generates a professional PDF report with formatted tables.

## Output Structure

The generated reports include:

### Addon Providers
- Summary table with provider IDs, names, and plan counts
- Detailed plans table with IDs, names, and slugs
- Grouped sections by provider with complete plan listings

### Application Instances
- Summary table with instance types, names, versions, and flavor counts
- Detailed flavors table with memory, CPU, pricing, and feature flags
- Grouped sections by application type with complete specifications

## Development

### Project Structure

```
cc-plans-lister/
├── cmd/cc-plans-lister/    # Main application entry point
├── internal/               # Private application code
│   ├── api/               # Clever Cloud API client
│   ├── config/            # Configuration management
│   └── formatters/        # Output format implementations
├── pkg/clevercloud/       # Public types and interfaces
├── test/                  # Test files and fixtures
├── go.mod                 # Go module definition
├── go.sum                 # Go module checksums
├── README.md              # This file
├── CLAUDE.md              # Claude Code guidance
└── .gitignore            # Git ignore rules
```

### Building

```bash
# Build for current platform
make build

# Build for multiple platforms (Linux, macOS Intel/ARM, Windows)
make build-all

# Clean build artifacts
make clean
```

### Running tests

```bash
# Run all tests
make test

# Run tests with coverage report (generates coverage.html)
make coverage

# Run tests with verbose output
go test -v ./...
```

### Code quality

```bash
# Format code
make fmt

# Run static analysis
make vet

# Complete development workflow (format + vet + test + build)
make dev

# View all available commands
make help
```

## Dependencies

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [gofpdf](https://github.com/jung-kurt/gofpdf) - PDF generation
- [testify](https://github.com/stretchr/testify) - Testing toolkit
- [Clever Cloud Go Client](https://go.clever-cloud.dev/client) - Official API client

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

Please ensure that:
- All tests pass
- Code is properly formatted (`go fmt`)
- New features include appropriate tests
- Documentation is updated as needed

## License

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details.

This is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

## Troubleshooting

### Common Issues

**Authentication Error**
```
Error: failed to load configuration: CLEVER_API_TOKEN environment variable is required
```
Solution: Set the `CLEVER_API_TOKEN` environment variable with your API token.

**API Connection Issues**
```
Error: failed to fetch addon providers: network error
```
Solution: Check your internet connection and verify that the API token is valid.

**Invalid Output Format**
```
Error: unsupported output format: xyz (supported: markdown, txt, csv, pdf)
```
Solution: Use one of the supported formats: `markdown`, `txt`, `csv`, or `pdf`.

### Getting Help

- Check the built-in help: `./bin/cc-plans-lister --help`
- View version information: `./bin/cc-plans-lister version`
- Review the examples in this README

## Examples

### Generate all formats

```bash
# Set your API token
export CLEVER_API_TOKEN="your_token_here"

# Build and run with Makefile
make build
./bin/cc-plans-lister --format=markdown --output=services.md
./bin/cc-plans-lister --format=txt --output=services.txt
./bin/cc-plans-lister --format=csv --output=services.csv
./bin/cc-plans-lister --format=pdf --output=services.pdf

# Or use the samples target to generate all formats at once
make samples  # Generates sample.{md,txt,csv,pdf}
```

### Automation example

```bash
#!/bin/bash
# Script to generate daily reports

DATE=$(date +%Y%m%d)
OUTPUT_DIR="reports/$DATE"

# Build the application
make build

# Create output directory
mkdir -p "$OUTPUT_DIR"

# Generate reports
./bin/cc-plans-lister --format=markdown --output="$OUTPUT_DIR/services.md"
./bin/cc-plans-lister --format=csv --output="$OUTPUT_DIR/services.csv" 
./bin/cc-plans-lister --format=pdf --output="$OUTPUT_DIR/services.pdf"

echo "Reports generated in $OUTPUT_DIR"

# Optional: Clean up previous day's sample files
make clean
```