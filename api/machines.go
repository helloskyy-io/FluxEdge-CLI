// api/machines.go
package api

import (
	"encoding/json"
	"fmt"
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

// GetMachines fetches available machines
func (c *Client) GetMachines() ([]Machine, error) {
	body, err := c.request("GET", "/computers")
	if err != nil {
		return nil, fmt.Errorf("API request failed: %w", err)
	}

	// Adjust parsing to match JSON structure
	var response MachinesResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse API response: %w", err)
	}

	return response.Computers, nil
}

