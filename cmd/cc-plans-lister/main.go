package main

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"cc-plans-lister/internal/api"
	"cc-plans-lister/internal/config"
	"cc-plans-lister/internal/formatters"
)

var (
	outputFormat string
	outputFile   string
	version      = "1.0.0"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cc-plans-lister",
	Short: "List Clever Cloud addon providers and application instances with their plans/flavors",
	Long: `cc-plans-lister fetches data from the Clever Cloud API to generate comprehensive 
documentation of available addon providers and application instance types with their 
respective plans and flavors.

The tool supports multiple output formats: markdown, txt, csv, and pdf.

Authentication is required via the CLEVER_API_TOKEN environment variable.`,
	RunE: runList,
}

func init() {
	rootCmd.Flags().StringVarP(&outputFormat, "format", "f", "markdown", "Output format (markdown, txt, csv, pdf)")
	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file (default: stdout)")

	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("cc-plans-lister v%s\n", version)
	},
}

func runList(cmd *cobra.Command, args []string) error {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	// Validate output format
	if !config.ValidateOutputFormat(outputFormat) {
		return fmt.Errorf("unsupported output format: %s (supported: markdown, txt, csv, pdf)", outputFormat)
	}

	// Create API client
	client := api.NewClient(cfg.APIToken)

	ctx := context.Background()

	// Fetch addon providers
	fmt.Fprintln(os.Stderr, "Fetching addon providers from Clever Cloud API...")
	providers, err := client.GetAddonProviders(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch addon providers: %w", err)
	}

	// Fetch product instances
	fmt.Fprintln(os.Stderr, "Fetching application instances from Clever Cloud API...")
	instances, err := client.GetProductInstances(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch product instances: %w", err)
	}

	// Get formatter
	formatter := formatters.GetFormatter(outputFormat)

	// Determine output destination
	var output *os.File
	if outputFile == "" {
		output = os.Stdout
	} else {
		output, err = os.Create(outputFile)
		if err != nil {
			return fmt.Errorf("failed to create output file: %w", err)
		}
		defer output.Close()
	}

	// Generate output
	err = formatter.Format(providers, instances, output)
	if err != nil {
		return fmt.Errorf("failed to format output: %w", err)
	}

	// Success message (only if writing to file)
	if outputFile != "" {
		fmt.Fprintf(os.Stderr, "Successfully generated %s with %d addon providers and %d application types\n",
			outputFile, len(providers), len(instances))
	}

	return nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
