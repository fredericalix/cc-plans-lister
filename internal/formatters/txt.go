package formatters

import (
	"fmt"
	"io"
	"sort"
	"strings"
	"text/tabwriter"

	"cc-plans-lister/pkg/clevercloud"
)

// TextFormatter generates plain text tabular output
type TextFormatter struct{}

// Format generates plain text tabular output for addon providers and product instances
func (f *TextFormatter) Format(providers []clevercloud.AddonProvider, instances []clevercloud.ProductInstance, writer io.Writer) error {
	var builder strings.Builder
	
	builder.WriteString("COMPLETE CLEVER CLOUD SERVICES OVERVIEW\n")
	builder.WriteString("========================================\n\n")
	builder.WriteString("Automatically generated via Clever Cloud API\n\n")
	
	// Addon Summary
	builder.WriteString("ADDON SUMMARY\n")
	builder.WriteString("=============\n\n")
	
	w := tabwriter.NewWriter(&builder, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "Provider ID\tName\tNumber of Plans")
	fmt.Fprintln(w, "-----------\t----\t---------------")
	
	for _, provider := range providers {
		fmt.Fprintf(w, "%s\t%s\t%d\n", provider.ID, provider.Name, len(provider.Plans))
	}
	w.Flush()
	
	// Application Summary
	builder.WriteString("\n\nAPPLICATION SUMMARY\n")
	builder.WriteString("===================\n\n")
	
	w = tabwriter.NewWriter(&builder, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "Type\tName\tVersion\tEnabled\tFlavors\tDefault Flavor")
	fmt.Fprintln(w, "----\t----\t-------\t-------\t-------\t--------------")
	
	for _, instance := range instances {
		enabledStr := "No"
		if instance.Enabled {
			enabledStr = "Yes"
		}
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%d\t%s\n",
			instance.Type, instance.Name, instance.Version, enabledStr, len(instance.Flavors), instance.DefaultFlavor.Name)
	}
	w.Flush()

	// Detailed Addon Plans
	builder.WriteString("\n\nDETAILED ADDON PLANS\n")
	builder.WriteString("====================\n\n")
	
	w = tabwriter.NewWriter(&builder, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "Provider ID\tProvider Name\tPlan ID\tPlan Name\tPlan Slug")
	fmt.Fprintln(w, "-----------\t-------------\t-------\t---------\t---------")
	
	for _, provider := range providers {
		if len(provider.Plans) == 0 {
			fmt.Fprintf(w, "%s\t%s\t-\tNo plans available\t-\n", provider.ID, provider.Name)
			continue
		}
		
		// Sort plans by slug for consistent output
		plans := make([]clevercloud.AddonPlan, len(provider.Plans))
		copy(plans, provider.Plans)
		sort.Slice(plans, func(i, j int) bool {
			return plans[i].Slug < plans[j].Slug
		})
		
		for i, plan := range plans {
			providerID := ""
			providerName := ""
			if i == 0 {
				providerID = provider.ID
				providerName = provider.Name
			}
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n",
				providerID, providerName, plan.ID, plan.Name, plan.Slug)
		}
	}
	w.Flush()

	// Detailed Application Flavors
	builder.WriteString("\n\nDETAILED APPLICATION FLAVORS\n")
	builder.WriteString("============================\n\n")
	
	w = tabwriter.NewWriter(&builder, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "Type\tName\tFlavor\tMemory\tCPU\tPrice\tAvailable\tMicroservice\tML")
	fmt.Fprintln(w, "----\t----\t------\t------\t---\t-----\t---------\t------------\t--")
	
	for _, instance := range instances {
		if !instance.Enabled {
			continue
		}
		
		if len(instance.Flavors) == 0 {
			fmt.Fprintf(w, "%s\t%s\t-\t-\t-\t-\t-\t-\t-\n", instance.Type, instance.Name)
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
				typeCell = instance.Type
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
			
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%d\t%.2f€\t%s\t%s\t%s\n",
				typeCell, nameCell, flavor.Name, flavor.Memory.Formatted, flavor.Cpus, flavor.Price, availableStr, microserviceStr, mlStr)
		}
	}
	w.Flush()

	// Plans by Provider
	builder.WriteString("\n\nPLANS BY ADDON PROVIDER\n")
	builder.WriteString("=======================\n\n")
	
	for _, provider := range providers {
		builder.WriteString(fmt.Sprintf("%s (%s)\n", provider.Name, provider.ID))
		builder.WriteString(strings.Repeat("-", len(provider.Name)+len(provider.ID)+3) + "\n")
		
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
			builder.WriteString(fmt.Sprintf("- %s (%s) - ID: %s\n", plan.Name, plan.Slug, plan.ID))
		}
		builder.WriteString("\n")
	}

	// Flavors by Application Type
	builder.WriteString("FLAVORS BY APPLICATION TYPE\n")
	builder.WriteString("===========================\n\n")
	
	for _, instance := range instances {
		if !instance.Enabled {
			continue
		}
		
		title := fmt.Sprintf("%s (%s) - Version %s", instance.Name, instance.Type, instance.Version)
		builder.WriteString(title + "\n")
		builder.WriteString(strings.Repeat("-", len(title)) + "\n")
		builder.WriteString(fmt.Sprintf("Description: %s\n", instance.Description))
		builder.WriteString(fmt.Sprintf("Max instances: %d\n", instance.MaxInstances))
		builder.WriteString(fmt.Sprintf("Tags: %s\n", strings.Join(instance.Tags, ", ")))
		builder.WriteString(fmt.Sprintf("Deployments: %s\n\n", strings.Join(instance.Deployments, ", ")))
		
		if len(instance.Flavors) == 0 {
			builder.WriteString("No flavors available.\n\n")
			continue
		}
		
		flavors := make([]clevercloud.Flavor, len(instance.Flavors))
		copy(flavors, instance.Flavors)
		sort.Slice(flavors, func(i, j int) bool {
			return flavors[i].Name < flavors[j].Name
		})
		
		builder.WriteString("Available flavors:\n")
		for _, flavor := range flavors {
			defaultMarker := ""
			if flavor.Name == instance.DefaultFlavor.Name {
				defaultMarker = " (default)"
			}
			
			builder.WriteString(fmt.Sprintf("- %s%s - %s, %d CPU, %.2f€/h",
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
				builder.WriteString(fmt.Sprintf(" [%s]", strings.Join(tags, ", ")))
			}
			builder.WriteString("\n")
		}
		builder.WriteString("\n")
	}
	
	_, err := writer.Write([]byte(builder.String()))
	return err
}