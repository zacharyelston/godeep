// internal/client/dataset.go
package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (c *DeepLakeClient) CreateDataset() error {
	fmt.Println("Creating tensor_db enabled dataset...")

	// Parse the dataset path and remove hub:// prefix
	datasetPath := strings.TrimPrefix(c.config.Activeloop.DatasetPath, "hub://")
	parts := strings.Split(datasetPath, "/")
	if len(parts) != 2 {
		return fmt.Errorf("invalid dataset path format: %s", datasetPath)
	}

	// First verify if dataset exists
	exists, err := c.verifyDatasetExists()
	if err != nil {
		return fmt.Errorf("error checking dataset existence: %v", err)
	}

	if exists {
		fmt.Printf("Dataset already exists: %s\n", c.config.Activeloop.DatasetPath)
		return nil
	}

	// Construct the URL with organization and dataset name
	url := fmt.Sprintf("%s/api/datasets/v1/%s/%s", c.config.Activeloop.BaseURL, parts[0], parts[1])

	// Create request body with minimal required fields
	createReq := struct {
		TensorDB bool                   `json:"tensor_db"`
		Schema   map[string]interface{} `json:"schema"`
	}{
		TensorDB: true,
		Schema:   c.config.Activeloop.DefaultSchema,
	}

	body, err := json.Marshal(createReq)
	if err != nil {
		return fmt.Errorf("error marshaling request: %v", err)
	}

	// Create request with PUT method
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	// Set headers
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "DeepLake-Go-Client/1.0")
	req.Header.Set("X-Organization-ID", c.config.Activeloop.OrgID)

	fmt.Printf("\nMaking dataset creation request to: %s\n", url)
	fmt.Printf("Request headers: %v\n", req.Header)
	fmt.Printf("Request body: %s\n", string(body))

	// Make request
	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("dataset creation failed: %v", err)
	}
	defer resp.Body.Close()

	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %v", err)
	}

	fmt.Printf("Response status: %d\n", resp.StatusCode)
	fmt.Printf("Response headers: %v\n", resp.Header)
	fmt.Printf("Response body: %s\n", string(respBody))

	if resp.StatusCode >= 400 {
		var errorResp struct {
			Detail      string `json:"detail"`
			Description string `json:"description"`
			Message     string `json:"message"`
		}
		if err := json.Unmarshal(respBody, &errorResp); err == nil {
			if errorResp.Detail != "" {
				return fmt.Errorf("dataset creation failed: %s", errorResp.Detail)
			}
			if errorResp.Description != "" {
				return fmt.Errorf("dataset creation failed: %s", errorResp.Description)
			}
			if errorResp.Message != "" {
				return fmt.Errorf("dataset creation failed: %s", errorResp.Message)
			}
		}
		return fmt.Errorf("dataset creation failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	// Final verification
	exists, err = c.verifyDatasetExists()
	if err != nil {
		return fmt.Errorf("error verifying dataset creation: %v", err)
	}
	if !exists {
		return fmt.Errorf("dataset creation failed: dataset does not exist after creation")
	}

	fmt.Printf("Dataset created successfully: %s\n", c.config.Activeloop.DatasetPath)
	return nil
}

func (c *DeepLakeClient) verifyDatasetExists() (bool, error) {
	fmt.Println("Checking dataset existence...")

	datasetPath := strings.TrimPrefix(c.config.Activeloop.DatasetPath, "hub://")
	parts := strings.Split(datasetPath, "/")

	url := fmt.Sprintf("%s/api/datasets/v1/%s/%s", c.config.Activeloop.BaseURL, parts[0], parts[1])

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, fmt.Errorf("error creating verification request: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "DeepLake-Go-Client/1.0")
	req.Header.Set("X-Organization-ID", c.config.Activeloop.OrgID)

	resp, err := c.client.Do(req)
	if err != nil {
		return false, fmt.Errorf("verification request failed: %v", err)
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK, nil
}
