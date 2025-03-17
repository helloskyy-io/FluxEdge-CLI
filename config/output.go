// /config/output.go
package config

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"github.com/olekukonko/tablewriter"
)

// PrintOutput handles all CLI output
// - JSON mode outputs structured JSON
// - Default mode formats structured data as tables
func PrintOutput(data interface{}, msgType string, formatter func(interface{}) string) {
	// ✅ If JSON mode is enabled, structure output accordingly
	if OutputFormat == "json" || os.Getenv("OUTPUT_JSON") == "true" {
		output := map[string]interface{}{}

		switch msgType {
		case "log":
			output["log"] = data
		case "error":
			output["error"] = data
		case "data":
			output["data"] = data
		default:
			output["message"] = data
		}

		// Convert to JSON and print
		jsonOutput, err := json.MarshalIndent(output, "", "  ")
		if err != nil {
			fmt.Println(`{"error": "Failed to encode output to JSON"}`)
			return
		}

		fmt.Println(string(jsonOutput))
		return
	}

	// ✅ If JSON is NOT enabled, always use table formatting for structured data
	if msgType == "log" {
		fmt.Println(data) // Print standard logs
	} else if msgType == "error" {
		fmt.Println("❌ ERROR:", data) // Print error messages
	} else {
		// ✅ Format all structured data as tables (even single-value results)
		printTable(data)
	}
}

// printTable formats structured data into a table
func printTable(data interface{}) {
	// ✅ Ensure data is a slice of structs
	slice := reflect.ValueOf(data)

	// Handle single object case by converting to a slice
	if slice.Kind() != reflect.Slice {
		slice = reflect.ValueOf([]interface{}{data})
	}

	// ✅ Ensure slice is not empty
	if slice.Len() == 0 {
		fmt.Println("⚠️ No available data.")
		return
	}

	// ✅ Extract struct fields dynamically
	firstItem := slice.Index(0)
	typeOfItem := firstItem.Type()
	headers := []string{}
	for i := 0; i < typeOfItem.NumField(); i++ {
		headers = append(headers, typeOfItem.Field(i).Name)
	}

	// ✅ Initialize Table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)

	// ✅ Populate Table Rows
	for i := 0; i < slice.Len(); i++ {
		row := []string{}
		item := slice.Index(i)
		for j := 0; j < item.NumField(); j++ {
			row = append(row, fmt.Sprintf("%v", item.Field(j).Interface()))
		}
		table.Append(row)
	}

	// ✅ Render Table
	table.Render()
}

