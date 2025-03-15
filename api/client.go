package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client struct for API interaction
type Client struct {
	APIKey string
}

// NewClient creates a new API client
func NewClient(apiKey string) *Client {
	return &Client{APIKey: apiKey}
}

// Generic API request function
func (c *Client) request(method, endpoint string) ([]byte, error) {
	url := fmt.Sprintf("https://pouwtest.fluxcore.ai/api/v1%s", endpoint)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
