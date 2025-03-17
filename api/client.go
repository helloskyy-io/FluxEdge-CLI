// api/client.go
package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/helloskyy-io/FluxEdge-CLI/config"
)

// Client struct for API interaction
type Client struct {
	APIKey string
	APIURL string
}

// NewClient creates a new API client
func NewClient(apiKey string) *Client {
	cfg := config.LoadConfig()
	apiKey = config.GetAPIKey()
	return &Client{APIKey: apiKey, APIURL: cfg.APIURL}
}

// Generic API request function with centralized error handling
func (c *Client) request(method, endpoint string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", c.APIURL, endpoint)
	if config.DebugMode {
		config.PrintOutput(fmt.Sprintf("DEBUG: Sending request -> %s %s", method, url), "log", nil)
	}

	// ✅ Create HTTP request
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		config.PrintOutput(fmt.Sprintf("failed to create request: %v", err), "error", nil)
		return nil, err
	}

	// ✅ Set headers
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		config.PrintOutput(fmt.Sprintf("failed to send request: %v", err), "error", nil)
		return nil, err
	}
	defer resp.Body.Close()

	// ✅ Handle non-200 responses
	if resp.StatusCode != http.StatusOK {
		errorMsg := fmt.Sprintf("API error %d: %s", resp.StatusCode, http.StatusText(resp.StatusCode))
		config.PrintOutput(errorMsg, "error", nil)
		return nil, fmt.Errorf(errorMsg)
	}

	// ✅ Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		config.PrintOutput(fmt.Sprintf("failed to read response body: %v", err), "error", nil)
		return nil, err
	}

	if config.DebugMode {
		config.PrintOutput(fmt.Sprintf("DEBUG: API Response -> %s", string(body)), "log", nil)
	}

	return body, nil
}
