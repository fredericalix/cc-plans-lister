package formatters

import (
	"io"

	"cc-plans-lister/pkg/clevercloud"
)

// Formatter defines the interface for output formatters
type Formatter interface {
	Format(providers []clevercloud.AddonProvider, instances []clevercloud.ProductInstance, writer io.Writer) error
}

// GetFormatter returns the appropriate formatter based on the format string
func GetFormatter(format string) Formatter {
	switch format {
	case "markdown":
		return &MarkdownFormatter{}
	case "txt":
		return &TextFormatter{}
	case "csv":
		return &CSVFormatter{}
	case "pdf":
		return &PDFFormatter{}
	default:
		return &MarkdownFormatter{} // default to markdown
	}
}
