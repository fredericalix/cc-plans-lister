package formatters

import (
	"encoding/csv"
	"io"
	"sort"
	"strconv"
	"strings"

	"cc-plans-lister/pkg/clevercloud"
)

// CSVFormatter generates CSV output
type CSVFormatter struct{}

// Format generates CSV output for addon providers and product instances
func (f *CSVFormatter) Format(providers []clevercloud.AddonProvider, instances []clevercloud.ProductInstance, writer io.Writer) error {
	csvWriter := csv.NewWriter(writer)
	defer csvWriter.Flush()

	// Write header comment (as a single-column row)
	err := csvWriter.Write([]string{"# Complete Clever Cloud Services Overview - CSV Export"})
	if err != nil {
		return err
	}

	err = csvWriter.Write([]string{"# Automatically generated via Clever Cloud API"})
	if err != nil {
		return err
	}

	err = csvWriter.Write([]string{}) // empty row
	if err != nil {
		return err
	}

	// Addon Providers section
	err = csvWriter.Write([]string{"# ADDON PROVIDERS"})
	if err != nil {
		return err
	}

	// Addon providers header
	err = csvWriter.Write([]string{
		"Type", "Provider_ID", "Provider_Name", "Plan_ID", "Plan_Name", "Plan_Slug",
	})
	if err != nil {
		return err
	}

	// Write addon data
	for _, provider := range providers {
		if len(provider.Plans) == 0 {
			err = csvWriter.Write([]string{
				"addon", provider.ID, provider.Name, "", "No plans available", "",
			})
			if err != nil {
				return err
			}
			continue
		}

		// Sort plans by slug for consistent output
		plans := make([]clevercloud.AddonPlan, len(provider.Plans))
		copy(plans, provider.Plans)
		sort.Slice(plans, func(i, j int) bool {
			return plans[i].Slug < plans[j].Slug
		})

		for _, plan := range plans {
			err = csvWriter.Write([]string{
				"addon",
				provider.ID,
				provider.Name,
				plan.ID,
				plan.Name,
				plan.Slug,
			})
			if err != nil {
				return err
			}
		}
	}

	// Empty row separator
	err = csvWriter.Write([]string{})
	if err != nil {
		return err
	}

	// Application Instances section
	err = csvWriter.Write([]string{"# APPLICATION INSTANCES"})
	if err != nil {
		return err
	}

	// Application instances header
	err = csvWriter.Write([]string{
		"Type", "Instance_Type", "Instance_Name", "Version", "Description", "Enabled", "Max_Instances",
		"Tags", "Deployments", "Flavor_Name", "Memory_Formatted", "Memory_Value", "Memory_Unit",
		"CPUs", "GPUs", "Price", "Available", "Microservice", "MachineLearning", "IsDefault",
	})
	if err != nil {
		return err
	}

	// Write application data
	for _, instance := range instances {
		if len(instance.Flavors) == 0 {
			err = csvWriter.Write([]string{
				"application",
				instance.Type,
				instance.Name,
				instance.Version,
				instance.Description,
				strconv.FormatBool(instance.Enabled),
				strconv.Itoa(instance.MaxInstances),
				strings.Join(instance.Tags, "|"),
				strings.Join(instance.Deployments, "|"),
				"", "", "", "", "", "", "", "", "", "", "",
			})
			if err != nil {
				return err
			}
			continue
		}

		// Sort flavors by name for consistent output
		flavors := make([]clevercloud.Flavor, len(instance.Flavors))
		copy(flavors, instance.Flavors)
		sort.Slice(flavors, func(i, j int) bool {
			return flavors[i].Name < flavors[j].Name
		})

		for _, flavor := range flavors {
			isDefault := flavor.Name == instance.DefaultFlavor.Name

			err = csvWriter.Write([]string{
				"application",
				instance.Type,
				instance.Name,
				instance.Version,
				instance.Description,
				strconv.FormatBool(instance.Enabled),
				strconv.Itoa(instance.MaxInstances),
				strings.Join(instance.Tags, "|"),
				strings.Join(instance.Deployments, "|"),
				flavor.Name,
				flavor.Memory.Formatted,
				strconv.Itoa(flavor.Memory.Value),
				flavor.Memory.Unit,
				strconv.Itoa(flavor.Cpus),
				strconv.Itoa(flavor.Gpus),
				strconv.FormatFloat(flavor.Price, 'f', 2, 64),
				strconv.FormatBool(flavor.Available),
				strconv.FormatBool(flavor.Microservice),
				strconv.FormatBool(flavor.MachineLearning),
				strconv.FormatBool(isDefault),
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}
