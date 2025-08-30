package formatters

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"cc-plans-lister/pkg/clevercloud"
)

// MarkdownFormatter generates markdown output
type MarkdownFormatter struct{}

// Format generates a complete markdown table for addon providers and product instances
func (f *MarkdownFormatter) Format(providers []clevercloud.AddonProvider, instances []clevercloud.ProductInstance, writer io.Writer) error {
	var builder strings.Builder
	
	builder.WriteString("# Complete Clever Cloud Services Overview\n\n")
	builder.WriteString("This document lists all available addon types AND application types on Clever Cloud with their respective plans/flavors.\n\n")
	builder.WriteString("*Automatically generated via Clever Cloud API*\n\n")
	
	// Summary table for addons
	builder.WriteString("## Addon Summary\n\n")
	builder.WriteString("| Provider ID | Name | Number of Plans |\n")
	builder.WriteString("|-------------|------|----------------|\n")
	
	for _, provider := range providers {
		builder.WriteString(fmt.Sprintf("| `%s` | %s | %d |\n", 
			provider.ID, provider.Name, len(provider.Plans)))
	}

	// Summary table for applications
	builder.WriteString("\n## Application Summary\n\n")
	builder.WriteString("| Type | Name | Version | Enabled | Number of Flavors | Default Flavor |\n")
	builder.WriteString("|------|------|---------|---------|-------------------|----------------|\n")
	
	for _, instance := range instances {
		enabledStr := "No"
		if instance.Enabled {
			enabledStr = "Yes"
		}
		builder.WriteString(fmt.Sprintf("| `%s` | %s | %s | %s | %d | `%s` |\n", 
			instance.Type, instance.Name, instance.Version, enabledStr, len(instance.Flavors), instance.DefaultFlavor.Name))
	}
	
	// Detailed addon plans table
	builder.WriteString("\n## Detailed Addon Plans\n\n")
	builder.WriteString("| Provider ID | Provider Name | Plan ID | Plan Name | Plan Slug |\n")
	builder.WriteString("|-------------|---------------|---------|-----------|----------|\n")
	
	for _, provider := range providers {
		if len(provider.Plans) == 0 {
			builder.WriteString(fmt.Sprintf("| `%s` | %s | - | No plans available | - |\n", 
				provider.ID, provider.Name))
			continue
		}
		
		// Sort plans by slug for consistent output
		plans := make([]clevercloud.AddonPlan, len(provider.Plans))
		copy(plans, provider.Plans)
		sort.Slice(plans, func(i, j int) bool {
			return plans[i].Slug < plans[j].Slug
		})
		
		for i, plan := range plans {
			providerCell := ""
			nameCell := ""
			if i == 0 {
				providerCell = fmt.Sprintf("`%s`", provider.ID)
				nameCell = provider.Name
			}
			
			builder.WriteString(fmt.Sprintf("| %s | %s | `%s` | %s | `%s` |\n",
				providerCell, nameCell, plan.ID, plan.Name, plan.Slug))
		}
	}

	// Detailed application flavors table
	builder.WriteString("\n## Detailed Application Flavors\n\n")
	builder.WriteString("| Type | Name | Flavor | Memory | CPU | Price | Available | Microservice | ML |\n")
	builder.WriteString("|------|------|--------|--------|-----|-------|-----------|-------------|----|\n")
	
	for _, instance := range instances {
		if !instance.Enabled {
			continue // Skip disabled instances
		}
		
		if len(instance.Flavors) == 0 {
			builder.WriteString(fmt.Sprintf("| `%s` | %s | - | - | - | - | - | - | - |\n", 
				instance.Type, instance.Name))
			continue
		}
		
		// Sort flavors by name for consistent output
		flavors := make([]clevercloud.Flavor, len(instance.Flavors))
		copy(flavors, instance.Flavors)
		sort.Slice(flavors, func(i, j int) bool {
			return flavors[i].Name < flavors[j].Name
		})
		
		for i, flavor := range flavors {
			typeCell := ""
			nameCell := ""
			if i == 0 {
				typeCell = fmt.Sprintf("`%s`", instance.Type)
				nameCell = instance.Name
			}
			
			availableStr := "No"
			if flavor.Available {
				availableStr = "Yes"
			}
			
			microserviceStr := "No"
			if flavor.Microservice {
				microserviceStr = "Yes"
			}
			
			mlStr := "No"
			if flavor.MachineLearning {
				mlStr = "Yes"
			}
			
			builder.WriteString(fmt.Sprintf("| %s | %s | `%s` | %s | %d | %.2f€ | %s | %s | %s |\n",
				typeCell, nameCell, flavor.Name, flavor.Memory.Formatted, flavor.Cpus, flavor.Price, availableStr, microserviceStr, mlStr))
		}
	}
	
	// Addons by provider section
	builder.WriteString("\n## Plans by Addon Provider\n\n")
	for _, provider := range providers {
		builder.WriteString(fmt.Sprintf("### %s (`%s`)\n\n", provider.Name, provider.ID))
		
		if len(provider.Plans) == 0 {
			builder.WriteString("No plans available.\n\n")
			continue
		}
		
		plans := make([]clevercloud.AddonPlan, len(provider.Plans))
		copy(plans, provider.Plans)
		sort.Slice(plans, func(i, j int) bool {
			return plans[i].Slug < plans[j].Slug
		})
		
		for _, plan := range plans {
			builder.WriteString(fmt.Sprintf("- **%s** (`%s`) - ID: `%s`\n", 
				plan.Name, plan.Slug, plan.ID))
		}
		builder.WriteString("\n")
	}

	// Applications by type section
	builder.WriteString("\n## Flavors by Application Type\n\n")
	for _, instance := range instances {
		if !instance.Enabled {
			continue // Skip disabled instances
		}
		
		builder.WriteString(fmt.Sprintf("### %s (`%s`) - Version %s\n\n", instance.Name, instance.Type, instance.Version))
		builder.WriteString(fmt.Sprintf("**Description**: %s\n\n", instance.Description))
		builder.WriteString(fmt.Sprintf("**Max instances**: %d\n\n", instance.MaxInstances))
		builder.WriteString(fmt.Sprintf("**Tags**: %s\n\n", strings.Join(instance.Tags, ", ")))
		builder.WriteString(fmt.Sprintf("**Deployments**: %s\n\n", strings.Join(instance.Deployments, ", ")))
		
		if len(instance.Flavors) == 0 {
			builder.WriteString("No flavors available.\n\n")
			continue
		}
		
		flavors := make([]clevercloud.Flavor, len(instance.Flavors))
		copy(flavors, instance.Flavors)
		sort.Slice(flavors, func(i, j int) bool {
			return flavors[i].Name < flavors[j].Name
		})
		
		builder.WriteString("**Available flavors**:\n\n")
		for _, flavor := range flavors {
			defaultMarker := ""
			if flavor.Name == instance.DefaultFlavor.Name {
				defaultMarker = " *(default)*"
			}
			
			builder.WriteString(fmt.Sprintf("- **%s**%s - %s, %d CPU, %.2f€/h", 
				flavor.Name, defaultMarker, flavor.Memory.Formatted, flavor.Cpus, flavor.Price))
			
			var tags []string
			if !flavor.Available {
				tags = append(tags, "Unavailable")
			}
			if flavor.Microservice {
				tags = append(tags, "Microservice")
			}
			if flavor.MachineLearning {
				tags = append(tags, "ML")
			}
			
			if len(tags) > 0 {
				builder.WriteString(fmt.Sprintf(" *[%s]*", strings.Join(tags, ", ")))
			}
			builder.WriteString("\n")
		}
		builder.WriteString("\n")
	}
	
	_, err := writer.Write([]byte(builder.String()))
	return err
}