package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/helloskyy-io/FluxEdge-CLI/config"
)

// API Base URL
const baseURL = "https://pouwtest.fluxcore.ai/api/v1"

// Client struct for API interaction
type Client struct {
	APIKey string
	APIURL string
}

// NewClient creates a new API client
func NewClient(apiKey string) *Client {
	cfg := config.LoadConfig()
	apiKey = config.LoadAPIKey()
	return &Client{APIKey: apiKey, APIURL: cfg.APIURL}
}

// Generic API request function
func (c *Client) request(method, endpoint string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", c.APIURL, endpoint)
	
	var req *http.Request
	var err error
	
	req, err = http.NewRequest(method, url, nil)
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

	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
