// config/config.go
package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Config struct for YAML settings
type Config struct {
	APIURL string `yaml:"api_url"`
	DebugMode bool   `yaml:"debug_mode"`
}

// Global variables
var DebugMode bool           		// Debug mode flag
var configLoaded bool        		// Ensures config loads once
var APILoadedFrom string     		// Tracks where API key was loaded from
var configData *Config       		// Store the config struct globally
var apiKey string                  	// Declare API Key globally
var OutputFormat string = "text"   	// ✅ Default output format

// LoadConfig reads config.yaml (Exits if missing/invalid)
func LoadConfig() *Config {
	if configLoaded {
		return configData
	}
	var config Config
	file, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		handleFatalError(fmt.Errorf("Critical Error: Missing config.yaml. This file is required to set the API endpoint."), 1)
		return nil
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		handleFatalError(fmt.Errorf("Critical Error: config.yaml is invalid. Please check formatting."), 1)
		return nil
	}

	// If the API URL is missing, that's also fatal
	if config.APIURL == "" {
		handleFatalError(fmt.Errorf("Critical Error: Missing API URL in config.yaml. This value is required."), 1)
		return nil
	}

	DebugMode = config.DebugMode
	configLoaded = true
	configData = &config

	return configData
}

// SetAPIKey stores the API key (called from cmd/root.go)
func SetAPIKey(key string, source string) {
	apiKey = key
	APILoadedFrom = source
}

// GetAPIKey retrieves the stored API key (used by api/client.go)
func GetAPIKey() string {
	if apiKey == "" {
		handleFatalError(fmt.Errorf("API_KEY is not set. Provide it via --api-key flag, environment variable, or .env file."), 3)
	}
	return apiKey
}

// LoadAPIKey initializes the API key following the correct priority order
func LoadAPIKey() {
	// ✅ First, check if API key was set via `SetAPIKey()` (from CLI flag)
	if apiKey != "" {
		return
	}

	// ✅ Second, try loading from environment variable
	apiKeyEnv := os.Getenv("API_KEY")
	if apiKeyEnv != "" {
		SetAPIKey(apiKeyEnv, "Environment Variable")
		return
	}

	// ✅ Third, try loading from `.env` file
	_ = godotenv.Load()
	apiKeyEnv = os.Getenv("API_KEY")
	if apiKeyEnv != "" {
		SetAPIKey(apiKeyEnv, ".env File")
		return
	}

	// ❌ Fatal error: No API key found
	handleFatalError(fmt.Errorf("API_KEY is not set. Provide it via --api-key flag, environment variable, or .env file."), 3)
}

// AddOutputFlag adds the `--output json` flag to a command
func AddOutputFlag(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&OutputFormat, "output", "o", "text", "Output format: text or json")
}

// handleFatalError manages fatal errors and formats them for JSON output if required
func handleFatalError(err error, exitCode int) {
	if os.Getenv("OUTPUT_JSON") == "true" {
		errorJSON, _ := json.MarshalIndent(map[string]interface{}{
			"error": err.Error(),
			"code":  exitCode,
		}, "", "  ")
		fmt.Println(string(errorJSON))
	} else {
		fmt.Println("❌ FATAL ERROR:", err)
	}

	os.Exit(exitCode)
}