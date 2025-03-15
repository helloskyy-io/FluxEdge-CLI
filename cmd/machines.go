package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/helloskyy-io/FluxEdge-CLI/api"
	"github.com/helloskyy-io/FluxEdge-CLI/config"
)

// "get-machines" command
var getMachinesCmd = &cobra.Command{
	Use:   "get-machines",
	Short: "Retrieve available machines",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := config.LoadAPIKey()
		client := api.NewClient(apiKey) 
		fmt.Println("Fetching available machines...")

		machines, err := client.GetMachines()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// ✅ Use the correct struct fields
		for _, machine := range machines {
			fmt.Printf("Cluster: %s | CPU: %s | Memory: %dMB | GPUs: %d | Price: $%.3f/hr | Region: %s | Storage: %dGB\n",
				machine.ClusterName, machine.CPU, machine.Memory, machine.GPUs, machine.PricePerHour, machine.Region, machine.Storage)
		}
	},
}

func init() {
	rootCmd.AddCommand(getMachinesCmd)
}
