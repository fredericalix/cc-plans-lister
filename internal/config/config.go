package config

import (
	"fmt"
	"os"
)

// Config holds application configuration
type Config struct {
	APIToken     string
	OutputFormat string
	OutputFile   string
}

// LoadConfig loads configuration from environment variables and command line flags
func LoadConfig() (*Config, error) {
	token := os.Getenv("CLEVER_API_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("CLEVER_API_TOKEN environment variable is required")
	}

	return &Config{
		APIToken:     token,
		OutputFormat: "markdown", // default format
		OutputFile:   "",         // default to stdout
	}, nil
}

// ValidateOutputFormat checks if the provided format is supported
func ValidateOutputFormat(format string) bool {
	validFormats := map[string]bool{
		"markdown": true,
		"txt":      true,
		"csv":      true,
		"pdf":      true,
	}
	return validFormats[format]
}
