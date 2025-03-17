// cmd/machines.go
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/helloskyy-io/FluxEdge-CLI/api"
	"github.com/helloskyy-io/FluxEdge-CLI/config"
)

var debugFlag bool
var outputFormat string

// FormatMachines converts machine data into a structured human-readable format
func FormatMachines(data interface{}) string {
	machines, ok := data.([]api.Machine)
	if !ok {
		return "❌ Error: Invalid data format received from API"
	}

	// ✅ Handle case where no machines are available
	if len(machines) == 0 {
		return "⚠️ No available machines found."
	}

	output := "Available Machines:\n"
	for _, machine := range machines {
		output += fmt.Sprintf(
			"Cluster: %s | CPU: %s | Memory: %dMB | GPUs: %d | Price: $%.3f/hr | Region: %s | Storage: %dGB\n",
			machine.ClusterName, machine.CPU, machine.Memory, machine.GPUs, machine.PricePerHour, machine.Region, machine.Storage)
	}
	return output
}

// "get-machines" command
var getMachinesCmd = &cobra.Command{
	Use:   "get-machines",
	Short: "Retrieve available machines",
	Run: func(cmd *cobra.Command, args []string) {
		config.DebugMode, _ = cmd.Flags().GetBool("debug") // Set global DebugMode

		// ✅ Log fetching message
		config.PrintOutput("Fetching available machines...", "log", nil)

		// ✅ Load API key and verify it's set
		config.LoadAPIKey()
		apiKey := config.GetAPIKey()
		if apiKey == "" {
			config.PrintOutput("❌ API key is missing. Use --api-key flag or set it in environment variables.", "error", nil)
			return
		}

		// ✅ Create API client
		client := api.NewClient(apiKey)

		// ✅ Fetch machines (Improved error handling)
		machines, err := client.GetMachines()
		if err != nil {
			// Classify errors
			if err.Error() == "unauthorized" {
				config.PrintOutput("❌ Unauthorized: Invalid API Key", "error", nil)
			} else if err.Error() == "timeout" {
				config.PrintOutput("⚠️ API request timed out. Please try again.", "error", nil)
			} else {
				config.PrintOutput(fmt.Sprintf("❌ API Error: %s", err.Error()), "error", nil)
			}
			return
		}

		// ✅ Handle "no machines available" case correctly
		if len(machines) == 0 {
			if config.OutputFormat == "json" {
				// ✅ Provide structured JSON response for automation
				emptyResponse := map[string]interface{}{
					"status":  "ok",
					"message": "No available machines found.",
					"data":    []api.Machine{}, // Empty array ensures valid JSON structure
				}
				config.PrintOutput(emptyResponse, "data", nil)
			} else {
				// ✅ Provide human-readable response
				config.PrintOutput("⚠️ No available machines found.", "log", nil)
			}
			return
		}

		// ✅ Print output (JSON or structured human-readable text)
		config.PrintOutput(machines, "data", FormatMachines)
	},
}

func init() {
	rootCmd.AddCommand(getMachinesCmd)
}
