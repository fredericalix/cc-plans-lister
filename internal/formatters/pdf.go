package formatters

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/jung-kurt/gofpdf"

	"cc-plans-lister/pkg/clevercloud"
)

// PDFFormatter generates PDF output
type PDFFormatter struct{}

// Format generates PDF output for addon providers and product instances
func (f *PDFFormatter) Format(providers []clevercloud.AddonProvider, instances []clevercloud.ProductInstance, writer io.Writer) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Set up fonts - gofpdf has built-in fonts
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(190, 10, "Complete Clever Cloud Services Overview")
	pdf.Ln(15)

	pdf.SetFont("Arial", "", 10)
	pdf.Cell(190, 5, "Automatically generated via Clever Cloud API")
	pdf.Ln(10)

	// Addon Summary Section
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(190, 10, "Addon Summary")
	pdf.Ln(12)

	pdf.SetFont("Arial", "B", 10)
	pdf.Cell(60, 8, "Provider ID")
	pdf.Cell(80, 8, "Name")
	pdf.Cell(30, 8, "Plans")
	pdf.Ln(8)

	pdf.SetFont("Arial", "", 9)
	for _, provider := range providers {
		pdf.Cell(60, 6, provider.ID)
		pdf.Cell(80, 6, truncateText(provider.Name, 35))
		pdf.Cell(30, 6, fmt.Sprintf("%d", len(provider.Plans)))
		pdf.Ln(6)
	}

	pdf.Ln(10)

	// Application Summary Section
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(190, 10, "Application Summary")
	pdf.Ln(12)

	pdf.SetFont("Arial", "B", 9)
	pdf.Cell(25, 8, "Type")
	pdf.Cell(45, 8, "Name")
	pdf.Cell(20, 8, "Version")
	pdf.Cell(20, 8, "Enabled")
	pdf.Cell(20, 8, "Flavors")
	pdf.Cell(30, 8, "Default")
	pdf.Ln(8)

	pdf.SetFont("Arial", "", 8)
	for _, instance := range instances {
		enabledStr := "No"
		if instance.Enabled {
			enabledStr = "Yes"
		}

		pdf.Cell(25, 6, truncateText(instance.Type, 12))
		pdf.Cell(45, 6, truncateText(instance.Name, 20))
		pdf.Cell(20, 6, truncateText(instance.Version, 10))
		pdf.Cell(20, 6, enabledStr)
		pdf.Cell(20, 6, fmt.Sprintf("%d", len(instance.Flavors)))
		pdf.Cell(30, 6, truncateText(instance.DefaultFlavor.Name, 15))
		pdf.Ln(6)
	}

	// Start new page for detailed information
	pdf.AddPage()

	// Detailed Addon Plans
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(190, 10, "Detailed Addon Plans")
	pdf.Ln(12)

	for _, provider := range providers {
		pdf.SetFont("Arial", "B", 11)
		pdf.Cell(190, 8, fmt.Sprintf("%s (%s)", provider.Name, provider.ID))
		pdf.Ln(8)

		pdf.SetFont("Arial", "", 9)
		if len(provider.Plans) == 0 {
			pdf.Cell(190, 6, "  No plans available")
			pdf.Ln(6)
			continue
		}

		// Sort plans by slug for consistent output
		plans := make([]clevercloud.AddonPlan, len(provider.Plans))
		copy(plans, provider.Plans)
		sort.Slice(plans, func(i, j int) bool {
			return plans[i].Slug < plans[j].Slug
		})

		for _, plan := range plans {
			text := fmt.Sprintf("  • %s (%s) - ID: %s", plan.Name, plan.Slug, plan.ID)
			pdf.Cell(190, 6, truncateText(text, 80))
			pdf.Ln(6)
		}
		pdf.Ln(4)
	}

	// Start new page for application flavors
	pdf.AddPage()

	// Detailed Application Flavors
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(190, 10, "Detailed Application Flavors")
	pdf.Ln(12)

	for _, instance := range instances {
		if !instance.Enabled {
			continue
		}

		pdf.SetFont("Arial", "B", 11)
		title := fmt.Sprintf("%s (%s) - Version %s", instance.Name, instance.Type, instance.Version)
		pdf.Cell(190, 8, title)
		pdf.Ln(8)

		pdf.SetFont("Arial", "", 9)
		pdf.Cell(190, 6, fmt.Sprintf("Description: %s", truncateText(instance.Description, 80)))
		pdf.Ln(6)
		pdf.Cell(190, 6, fmt.Sprintf("Max instances: %d", instance.MaxInstances))
		pdf.Ln(6)
		pdf.Cell(190, 6, fmt.Sprintf("Tags: %s", strings.Join(instance.Tags, ", ")))
		pdf.Ln(6)

		if len(instance.Flavors) == 0 {
			pdf.Cell(190, 6, "  No flavors available")
			pdf.Ln(6)
			continue
		}

		// Sort flavors by name for consistent output
		flavors := make([]clevercloud.Flavor, len(instance.Flavors))
		copy(flavors, instance.Flavors)
		sort.Slice(flavors, func(i, j int) bool {
			return flavors[i].Name < flavors[j].Name
		})

		pdf.Cell(190, 6, "Available flavors:")
		pdf.Ln(6)

		for _, flavor := range flavors {
			defaultMarker := ""
			if flavor.Name == instance.DefaultFlavor.Name {
				defaultMarker = " (default)"
			}

			availableStr := "Available"
			if !flavor.Available {
				availableStr = "Unavailable"
			}

			var features []string
			if flavor.Microservice {
				features = append(features, "Microservice")
			}
			if flavor.MachineLearning {
				features = append(features, "ML")
			}

			featuresStr := ""
			if len(features) > 0 {
				featuresStr = fmt.Sprintf(" [%s]", strings.Join(features, ", "))
			}

			text := fmt.Sprintf("  • %s%s - %s, %d CPU, %.2f€/h - %s%s",
				flavor.Name, defaultMarker, flavor.Memory.Formatted, flavor.Cpus,
				flavor.Price, availableStr, featuresStr)
			pdf.Cell(190, 6, truncateText(text, 90))
			pdf.Ln(6)
		}
		pdf.Ln(6)
	}

	// Write PDF to writer
	return pdf.Output(writer)
}

// truncateText truncates text to fit within specified length
func truncateText(text string, maxLen int) string {
	if maxLen <= 0 {
		return ""
	}
	if len(text) <= maxLen {
		return text
	}
	if maxLen <= 3 {
		if maxLen == 1 {
			return "."
		}
		return "..." // for maxLen 2 or 3
	}
	return text[:maxLen-3] + "..."
}