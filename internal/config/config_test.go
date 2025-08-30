package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name        string
		tokenValue  string
		expectError bool
	}{
		{
			name:        "valid token",
			tokenValue:  "valid_token_123",
			expectError: false,
		},
		{
			name:        "empty token",
			tokenValue:  "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up environment
			if tt.tokenValue != "" {
				os.Setenv("CLEVER_API_TOKEN", tt.tokenValue)
			} else {
				os.Unsetenv("CLEVER_API_TOKEN")
			}
			defer os.Unsetenv("CLEVER_API_TOKEN")

			// Test
			cfg, err := LoadConfig()

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, cfg)
			} else {
				require.NoError(t, err)
				require.NotNil(t, cfg)
				assert.Equal(t, tt.tokenValue, cfg.APIToken)
				assert.Equal(t, "markdown", cfg.OutputFormat)
				assert.Equal(t, "", cfg.OutputFile)
			}
		})
	}
}

func TestValidateOutputFormat(t *testing.T) {
	tests := []struct {
		format string
		valid  bool
	}{
		{"markdown", true},
		{"txt", true},
		{"csv", true},
		{"pdf", true},
		{"json", false},
		{"xml", false},
		{"", false},
		{"MARKDOWN", false}, // case sensitive
	}

	for _, tt := range tests {
		t.Run(tt.format, func(t *testing.T) {
			result := ValidateOutputFormat(tt.format)
			assert.Equal(t, tt.valid, result)
		})
	}
}