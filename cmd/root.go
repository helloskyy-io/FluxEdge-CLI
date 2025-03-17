// /cmd/root.go
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/helloskyy-io/FluxEdge-CLI/config"
)

// APIKey is a global variable for CLI input
var APIKey string

// Define root command
var rootCmd = &cobra.Command{
	Use:   "edgeapi",
	Short: "EdgeAPI CLI - Interact with the Edge platform API",
	Long:  "A command-line interface for managing Edge platform deployments",
	Run: func(cmd *cobra.Command, args []string) {
		config.PrintOutput("Welcome to Edge CLI! Use --help to see available commands.", "log", nil)
	},
}

// Execute runs the CLI
func Execute() {
	// Load config (ensures API URL is available)
	config.LoadConfig()

	// ✅ Set OutputFormat globally
	if os.Getenv("OUTPUT_JSON") == "true" || config.OutputFormat == "json" {
		config.OutputFormat = "json"
	}

	// ✅ If --api-key flag was passed, store it in config
	if APIKey != "" {
		config.PrintOutput("✅ API Key loaded from CLI flag", "log", nil)
		config.SetAPIKey(APIKey, "CLI Flag")
	} else {
		// ✅ Otherwise, load API key from environment variables or .env
		config.LoadAPIKey()
	}

	// ✅ Check if the DEBUG environment variable is set
	if os.Getenv("DEBUG") == "true" {
		config.DebugMode = true
		config.PrintOutput("DEBUG MODE ENABLED FROM ENV", "log", nil)
	}

	if err := rootCmd.Execute(); err != nil {
		config.PrintOutput(err.Error(), "error", nil) // ✅ Updated to match new function signature
		os.Exit(1)
	}
}

// Initialize CLI flags
func init() {
	// ✅ Use `config.AddOutputFlag()` to add the `--output json` flag
	config.AddOutputFlag(rootCmd)

	// Define a global debug flag (--debug)
	rootCmd.PersistentFlags().BoolVar(&config.DebugMode, "debug", false, "Enable debug mode")

	// Define global API key flag (--api-key)
	rootCmd.PersistentFlags().StringVar(&APIKey, "api-key", "", "API Key for authentication")
}