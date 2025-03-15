package config

import (
	"fmt"
	"os"
	"log"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Config struct for YAML settings
type Config struct {
	APIURL string `yaml:"api_url"`
}

// LoadConfig reads config.yaml
func LoadConfig() *Config {
	var config Config
	file, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		fmt.Println("❌ ERROR: Failed to read config file: config/config.yaml")
		log.Fatal(err) // Stop execution
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		fmt.Println("❌ ERROR: Failed to parse config.yaml. Check formatting.")
		log.Fatal(err) // Stop execution
	}

	return &config
}

// LoadAPIKey loads API key from environment variable
func LoadAPIKey() string {
	// ✅ Load .env file before reading environment variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println("⚠️ WARNING: No .env file found, using system environment variables.")
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Println("❌ ERROR: API_KEY is not set. Exiting.")
		log.Fatal("Missing API_KEY")
	}

	return apiKey
}