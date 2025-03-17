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
	APIURL    string `yaml:"api_url"`
	DebugMode bool   `yaml:"debug_mode"`
}

// Global variables
var (
	DebugMode       bool            // Debug mode flag
	configLoaded    bool            // Ensures config loads once
	APILoadedFrom   string          // Tracks where API key was loaded from
	configData      *Config         // Store the config struct globally
	apiKey          string          // API Key globally stored
	OutputFormat    string = "text" // Default output format (text/json)
)

// LoadConfig reads config.yaml (FATAL exit if missing/invalid)
func LoadConfig() *Config {
	if configLoaded {
		return configData
	}

	var config Config
	file, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		handleFatalError(fmt.Errorf("Critical Error: Missing config.yaml. This file is required to set the API endpoint."), 1)
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		handleFatalError(fmt.Errorf("Critical Error: config.yaml is invalid. Please check formatting."), 1)
	}

	// If API URL is missing, that’s fatal
	if config.APIURL == "" {
		handleFatalError(fmt.Errorf("Critical Error: Missing API URL in config.yaml. This value is required."), 1)
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

// LoadAPIKey initializes the API key using a strict priority order
func LoadAPIKey() {
	// ✅ 1st Priority: CLI Flag (`--api-key`)
	if apiKey != "" {
		return
	}

	// ✅ 2nd Priority: Environment Variable
	apiKeyEnv := os.Getenv("API_KEY")
	if apiKeyEnv != "" {
		SetAPIKey(apiKeyEnv, "Environment Variable")
		return
	}

	// ✅ 3rd Priority: `.env` File
	_ = godotenv.Load()
	apiKeyEnv = os.Getenv("API_KEY")
	if apiKeyEnv != "" {
		SetAPIKey(apiKeyEnv, ".env File")
		return
	}

	// ❌ All API key sources failed: Fatal Error
	handleFatalError(fmt.Errorf("API_KEY is missing. Provide it via --api-key flag, environment variable, or .env file."), 3)
}

// AddOutputFlag adds the `--output json` flag to a command
func AddOutputFlag(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&OutputFormat, "output", "o", "text", "Output format: text or json")
}

// handleFatalError logs fatal errors and exits the program
func handleFatalError(err error, exitCode int) {
	if OutputFormat == "json" || os.Getenv("OUTPUT_JSON") == "true" {
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

// handleError logs non-fatal errors (for debugging/logging purposes)
func handleError(err error, exitCode int, fatal bool) {
	if OutputFormat == "json" || os.Getenv("OUTPUT_JSON") == "true" {
		errorJSON, _ := json.MarshalIndent(map[string]interface{}{
			"error": err.Error(),
			"code":  exitCode,
		}, "", "  ")
		fmt.Println(string(errorJSON))
	} else {
		fmt.Println("❌ ERROR:", err)
	}

	// Only exit if marked as fatal
	if fatal {
		os.Exit(exitCode)
	}
}
