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
go mod download
go build -o bin/cc-plans-lister cmd/cc-plans-lister/main.go
```

### Install dependencies

```bash
go mod tidy
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
go build -o bin/cc-plans-lister cmd/cc-plans-lister/main.go

# Build for multiple platforms
GOOS=linux GOARCH=amd64 go build -o bin/cc-plans-lister-linux-amd64 cmd/cc-plans-lister/main.go
GOOS=darwin GOARCH=amd64 go build -o bin/cc-plans-lister-darwin-amd64 cmd/cc-plans-lister/main.go
GOOS=windows GOARCH=amd64 go build -o bin/cc-plans-lister-windows-amd64.exe cmd/cc-plans-lister/main.go
```

### Running tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...
```

### Code formatting

```bash
# Format all Go files
go fmt ./...

# Run linter (if installed)
golangci-lint run
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

This project is licensed under the MIT License - see the LICENSE file for details.

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

# Generate markdown report
./bin/cc-plans-lister --format=markdown --output=services.md

# Generate text report
./bin/cc-plans-lister --format=txt --output=services.txt

# Generate CSV export
./bin/cc-plans-lister --format=csv --output=services.csv

# Generate PDF report
./bin/cc-plans-lister --format=pdf --output=services.pdf
```

### Automation example

```bash
#!/bin/bash
# Script to generate daily reports

DATE=$(date +%Y%m%d)
OUTPUT_DIR="reports/$DATE"

mkdir -p "$OUTPUT_DIR"

./bin/cc-plans-lister --format=markdown --output="$OUTPUT_DIR/services.md"
./bin/cc-plans-lister --format=csv --output="$OUTPUT_DIR/services.csv"
./bin/cc-plans-lister --format=pdf --output="$OUTPUT_DIR/services.pdf"

echo "Reports generated in $OUTPUT_DIR"
```