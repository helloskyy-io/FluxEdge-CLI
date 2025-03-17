package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// PrintOutput handles all program output: logs, errors, and API responses
// formatter: A function that defines how human-readable output should be formatted.
func PrintOutput(data interface{}, msgType string, formatter func(interface{}) string) {
	// ✅ If JSON mode is enabled, structure output accordingly
	if OutputFormat == "json" || os.Getenv("OUTPUT_JSON") == "true" {
		output := map[string]interface{}{}

		// Classify messages
		if msgType == "log" {
			output["log"] = data
		} else if msgType == "error" {
			output["error"] = data
		} else if msgType == "data" {
			output["data"] = data
		}

		// Convert to JSON and print
		jsonOutput, _ := json.MarshalIndent(output, "", "  ")
		fmt.Println(string(jsonOutput))
		return
	}

	// ✅ If JSON is NOT enabled, use the formatter for human-readable output
	if msgType == "log" {
		fmt.Println(data) // Print standard logs
	} else if msgType == "error" {
		fmt.Println("❌ ERROR:", data) // Print error messages
	} else {
		// Use formatter for structured output
		fmt.Println(formatter(data))
	}
}
