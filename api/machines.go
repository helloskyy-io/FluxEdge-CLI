// api/machines.go
package api

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/helloskyy-io/FluxEdge-CLI/config"
)

// Machine struct matches the actual API response fields
type Machine struct {
	ClusterName  string  `json:"cluster_name"`
	CPU          string  `json:"cpu"`
	Hash         string  `json:"hash"`
	Memory       int     `json:"memory"`
	GPUs         int     `json:"nb_gpu"`
	PricePerHour float64 `json:"price_per_hour"`
	Region       string  `json:"region"`
	Storage      int     `json:"storage"`
}

// MachinesResponse wraps the API response
type MachinesResponse struct {
	Computers []Machine `json:"computers"`
}

// ExpectedFields contains the list of valid JSON keys for Machine struct
var ExpectedFields = map[string]bool{
	"cluster_name":  true,
	"cpu":           true,
	"hash":          true,
	"memory":        true,
	"nb_gpu":        true,
	"price_per_hour": true,
	"region":        true,
	"storage":       true,
}

// validateMachine ensures API response contains expected fields and no unexpected ones
func validateMachine(m Machine, rawMachine map[string]interface{}) (bool, []string, []string) {
	missingFields := []string{}
	unexpectedFields := []string{}

	// ✅ Check for missing required fields
	v := reflect.ValueOf(m)
	typeOfM := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := typeOfM.Field(i)
		jsonTag := field.Tag.Get("json")

		// If the JSON field is empty in the struct but expected, mark it as missing
		if v.Field(i).Interface() == reflect.Zero(field.Type).Interface() {
			missingFields = append(missingFields, jsonTag)
		}
	}

	// ✅ Check for unexpected fields (only inside each machine object)
	for key := range rawMachine {
		if _, expected := ExpectedFields[key]; !expected {
			unexpectedFields = append(unexpectedFields, key)
		}
	}

	return len(missingFields) == 0 && len(unexpectedFields) == 0, missingFields, unexpectedFields
}

// GetMachines fetches available machines and validates response format
func (c *Client) GetMachines() ([]Machine, error) {
	body, err := c.request("GET", "/computers")
	if err != nil {
		// ✅ Differentiate API failure causes
		apiError := fmt.Errorf("API request failed: %w", err)
		if config.DebugMode {
			config.PrintOutput(fmt.Sprintf("DEBUG: Raw API request error: %s", err.Error()), "log", nil)
		}
		return nil, apiError
	}

	// ✅ Log raw response in debug mode
	if config.DebugMode {
		config.PrintOutput(fmt.Sprintf("DEBUG: Raw API response: %s", string(body)), "log", nil)
	}

	// ✅ Parse JSON into struct
	var response MachinesResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		parseError := fmt.Errorf("failed to parse API response: %w", err)
		if config.DebugMode {
			config.PrintOutput(fmt.Sprintf("DEBUG: Failed response body: %s", string(body)), "log", nil)
		}
		return nil, parseError
	}

	// ✅ Handle case where no machines are available
	if len(response.Computers) == 0 {
		if config.DebugMode {
			config.PrintOutput("DEBUG: No available machines returned from API", "log", nil)
		}
		return []Machine{}, nil
	}

	// ✅ Validate response format
	validMachines := []Machine{}
	formatErrors := 0
	missingWarnings := 0
	extraWarnings := 0

	for _, machine := range response.Computers {
		// ✅ Parse only **this machine's JSON** into a map
		rawMachineData, err := json.Marshal(machine)
		if err != nil {
			config.PrintOutput("⚠️ ERROR: Unable to inspect machine object for validation.", "error", nil)
			continue
		}

		var rawMachine map[string]interface{}
		_ = json.Unmarshal(rawMachineData, &rawMachine)

		valid, missingFields, unexpectedFields := validateMachine(machine, rawMachine)
		if valid {
			validMachines = append(validMachines, machine)
		} else {
			formatErrors++

			// ✅ Log missing fields
			if len(missingFields) > 0 {
				missingWarnings++
				sort.Strings(missingFields)
				config.PrintOutput(fmt.Sprintf("⚠️ WARNING: Machine is missing expected fields: %s", strings.Join(missingFields, ", ")), "error", nil)
			}

			// ✅ Log unexpected fields
			if len(unexpectedFields) > 0 {
				extraWarnings++
				sort.Strings(unexpectedFields)
				config.PrintOutput(fmt.Sprintf("⚠️ WARNING: API returned unexpected fields: %s", strings.Join(unexpectedFields, ", ")), "error", nil)
			}
		}
	}

	// ✅ If some machines were filtered out due to errors, log a warning
	if formatErrors > 0 {
		if config.OutputFormat == "json" {
			warningJSON := map[string]interface{}{
				"warning":             "Some machines were ignored due to schema mismatches.",
				"invalid":             formatErrors,
				"missing_fields":      missingWarnings,
				"unexpected_fields":   extraWarnings,
				"total":               len(response.Computers),
				"valid":               len(validMachines),
			}
			config.PrintOutput(warningJSON, "error", nil)
		} else {
			config.PrintOutput(fmt.Sprintf("⚠️ WARNING: %d out of %d machines were ignored due to schema mismatches.", formatErrors, len(response.Computers)), "error", nil)
		}
	}

	return validMachines, nil
}
