package formatters

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"cc-plans-lister/test/fixtures"
)

func TestGetFormatter(t *testing.T) {
	tests := []struct {
		format   string
		expected interface{}
	}{
		{"markdown", &MarkdownFormatter{}},
		{"txt", &TextFormatter{}},
		{"csv", &CSVFormatter{}},
		{"pdf", &PDFFormatter{}},
		{"unknown", &MarkdownFormatter{}}, // default fallback
	}

	for _, tt := range tests {
		t.Run(tt.format, func(t *testing.T) {
			formatter := GetFormatter(tt.format)
			assert.IsType(t, tt.expected, formatter)
		})
	}
}

func TestMarkdownFormatter(t *testing.T) {
	formatter := &MarkdownFormatter{}
	var buf bytes.Buffer

	providers := fixtures.TestAddonProviders()
	instances := fixtures.TestProductInstances()

	err := formatter.Format(providers, instances, &buf)
	require.NoError(t, err)

	output := buf.String()

	// Check for expected headers
	assert.Contains(t, output, "# Complete Clever Cloud Services Overview")
	assert.Contains(t, output, "## Addon Summary")
	assert.Contains(t, output, "## Application Summary")
	assert.Contains(t, output, "## Detailed Addon Plans")
	assert.Contains(t, output, "## Detailed Application Flavors")

	// Check for test data
	assert.Contains(t, output, "Redis")
	assert.Contains(t, output, "PostgreSQL")
	assert.Contains(t, output, "Node.js")
	assert.Contains(t, output, "Python")
}

func TestTextFormatter(t *testing.T) {
	formatter := &TextFormatter{}
	var buf bytes.Buffer

	providers := fixtures.TestAddonProviders()
	instances := fixtures.TestProductInstances()

	err := formatter.Format(providers, instances, &buf)
	require.NoError(t, err)

	output := buf.String()

	// Check for expected headers
	assert.Contains(t, output, "COMPLETE CLEVER CLOUD SERVICES OVERVIEW")
	assert.Contains(t, output, "ADDON SUMMARY")
	assert.Contains(t, output, "APPLICATION SUMMARY")
	assert.Contains(t, output, "DETAILED ADDON PLANS")
	assert.Contains(t, output, "DETAILED APPLICATION FLAVORS")

	// Check for test data
	assert.Contains(t, output, "Redis")
	assert.Contains(t, output, "PostgreSQL")
	assert.Contains(t, output, "Node.js")
	assert.Contains(t, output, "Python")
}

func TestCSVFormatter(t *testing.T) {
	formatter := &CSVFormatter{}
	var buf bytes.Buffer

	providers := fixtures.TestAddonProviders()
	instances := fixtures.TestProductInstances()

	err := formatter.Format(providers, instances, &buf)
	require.NoError(t, err)

	output := buf.String()

	// Check for CSV structure
	assert.Contains(t, output, "# Complete Clever Cloud Services Overview - CSV Export")
	assert.Contains(t, output, "# ADDON PROVIDERS")
	assert.Contains(t, output, "# APPLICATION INSTANCES")
	assert.Contains(t, output, "Type,Provider_ID,Provider_Name")
	assert.Contains(t, output, "Type,Instance_Type,Instance_Name")

	// Check for test data
	assert.Contains(t, output, "redis,Redis")
	assert.Contains(t, output, "postgresql,PostgreSQL")
	assert.Contains(t, output, "node,Node.js")
	assert.Contains(t, output, "python,Python")
}

func TestPDFFormatter(t *testing.T) {
	formatter := &PDFFormatter{}
	var buf bytes.Buffer

	providers := fixtures.TestAddonProviders()
	instances := fixtures.TestProductInstances()

	err := formatter.Format(providers, instances, &buf)
	require.NoError(t, err)

	// Check that PDF content was generated (PDF starts with %PDF)
	output := buf.Bytes()
	assert.True(t, len(output) > 0, "PDF should generate content")
	assert.True(t, bytes.HasPrefix(output, []byte("%PDF")), "Output should be a valid PDF")
}

func TestTruncateText(t *testing.T) {
	tests := []struct {
		text     string
		maxLen   int
		expected string
	}{
		{"short", 10, "short"},
		{"this is a very long text", 10, "this is..."},
		{"exact", 5, "exact"},
		{"toolong", 5, "to..."},
		{"x", 3, "x"},
		{"toolong", 1, "."},
		{"toolong", 0, ""},
	}

	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			result := truncateText(tt.text, tt.maxLen)
			assert.Equal(t, tt.expected, result)
		})
	}
}