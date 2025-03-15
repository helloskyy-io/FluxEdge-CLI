package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// ✅ Define root command
var rootCmd = &cobra.Command{
	Use:   "edgeapi",
	Short: "EdgeAPI CLI - Interact with the Edge platform API",
	Long:  "A command-line interface for managing Edge platform deployments",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Edge CLI! Use --help to see available commands.")
	},
}

// Execute runs the CLI
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
