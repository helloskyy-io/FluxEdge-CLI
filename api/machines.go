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
		fmt.Println("API request failed:", err)
		return nil, err
	}

	// Debugging: Print raw response
	fmt.Println("Raw API Response:", string(body))

	// Adjust parsing to match JSON structure
	var response MachinesResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("API request failed:", err)
		return nil, err
	}

	return response.Computers, nil
}

// Test Function: Run this manually to check if API works
func TestGetMachines() {
	client := NewClient("521a8c01-bfe4-46a2-b335-0198cbe6eb57") // TEMP Key
	machines, err := client.GetMachines()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Machines List:")
	for _, machine := range machines {
		fmt.Printf("Cluster: %s | CPU: %s | Memory: %dMB | GPUs: %d | Price: $%.3f/hr | Region: %s | Storage: %dGB\n",
			machine.ClusterName, machine.CPU, machine.Memory, machine.GPUs, machine.PricePerHour, machine.Region, machine.Storage)
	}
}
