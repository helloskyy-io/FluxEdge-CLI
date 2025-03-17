// api/client.go
package api

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

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

// request makes an HTTP request with error handling and retries
func (c *Client) request(method, endpoint string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", c.APIURL, endpoint)
	if config.DebugMode {
		config.PrintOutput(fmt.Sprintf("DEBUG: Sending request -> %s %s", method, url), "log", nil)
	}

	// ✅ Create HTTP client with timeout
	httpClient := &http.Client{
		Timeout: 10 * time.Second, // Set timeout to avoid hanging requests
	}

	// ✅ Retry logic (up to 3 attempts)
	maxRetries := 3
	var lastErr error
	for attempt := 1; attempt <= maxRetries; attempt++ {
		// ✅ Create new request
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			config.PrintOutput(fmt.Sprintf("failed to create request: %v", err), "error", nil)
			return nil, err
		}

		// ✅ Set headers
		req.Header.Set("Authorization", "Bearer "+c.APIKey)
		req.Header.Set("Accept", "application/json")

		// ✅ Send request
		resp, err := httpClient.Do(req)
		if err != nil {
			// Handle timeout separately
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				config.PrintOutput(fmt.Sprintf("⚠️ API request timeout (Attempt %d/%d)", attempt, maxRetries), "error", nil)
			} else {
				config.PrintOutput(fmt.Sprintf("failed to send request: %v (Attempt %d/%d)", err, attempt, maxRetries), "error", nil)
			}

			lastErr = err
			time.Sleep(2 * time.Second) // Wait before retrying
			continue
		}
		defer resp.Body.Close()

		// ✅ Handle non-200 responses
		if resp.StatusCode != http.StatusOK {
			errorMsg := fmt.Sprintf("API error %d: %s", resp.StatusCode, http.StatusText(resp.StatusCode))

			// Specific error messages for common API issues
			switch resp.StatusCode {
			case http.StatusUnauthorized:
				errorMsg = "❌ Unauthorized: Invalid API Key."
			case http.StatusTooManyRequests:
				errorMsg = "⚠️ Rate Limit Exceeded: Too many requests. Try again later."
			case http.StatusInternalServerError, http.StatusBadGateway, http.StatusServiceUnavailable, http.StatusGatewayTimeout:
				errorMsg = fmt.Sprintf("⚠️ API temporarily unavailable (Attempt %d/%d)", attempt, maxRetries)
			}

			config.PrintOutput(errorMsg, "error", nil)
			lastErr = fmt.Errorf(errorMsg)

			// Retry only on server errors (500+)
			if resp.StatusCode >= 500 {
				time.Sleep(2 * time.Second) // Wait before retrying
				continue
			}

			return nil, lastErr
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

		return body, nil // Success
	}

	return nil, lastErr // Return last error after retries
}
