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
		return "Invalid data format"
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
		
		// API key 
		config.LoadAPIKey()
		apiKey := config.GetAPIKey()
		client := api.NewClient(apiKey)

		// Fetch machines
		machines, err := client.GetMachines()
		if err != nil {
			config.PrintOutput(err.Error(), "error", nil) // Handle errors through centralized function
			return
		}

		// Print output (JSON or structured human-readable text)
		config.PrintOutput(machines, "data", FormatMachines)
	},
}

func init() {
	rootCmd.AddCommand(getMachinesCmd)
}
