package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/DurgaPratapRajbhar/e-commerce/pkg/config"
)

// InventoryClient handles communication with the inventory service
type InventoryClient struct {
	baseURL string
	timeout time.Duration
}

// NewInventoryClient creates a new inventory service client
func NewInventoryClient(cfg *config.Config) *InventoryClient {
	return &InventoryClient{
		baseURL: cfg.Services.InventoryService.URL,
		timeout: cfg.Services.InventoryService.Timeout,
	}
}

// GetInventory fetches inventory with pagination from inventory service
func (c *InventoryClient) GetInventory(ctx context.Context, authToken string, page, limit string) ([]interface{}, error) {
	url := fmt.Sprintf("%s/inventory/items?page=%s&limit=%s", c.baseURL, page, limit)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	if authToken != "" {
		req.Header.Set("Authorization", authToken)
	}

	client := &http.Client{Timeout: c.timeout}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	// Check if it's an error response first
	if success, ok := response["success"].(bool); ok && !success {
		return nil, fmt.Errorf("inventory service returned error: %v", response["message"])
	}

	// Try to get the data field
	responseData, ok := response["data"]
	if !ok {
		// Check for alternative field names
		if dataAlt, okAlt := response["Data"].(map[string]interface{}); okAlt {
			// If data is directly an array
			if inventoryData, okArray := dataAlt["items"].([]interface{}); okArray {
				return inventoryData, nil
			}
			// If data has "Data" field
			if inventoryData, okArray := dataAlt["Data"].([]interface{}); okArray {
				return inventoryData, nil
			}
			// If data has "data" field
			if inventoryData, okArray := dataAlt["data"].([]interface{}); okArray {
				return inventoryData, nil
			}
			return nil, fmt.Errorf("invalid data format in inventory service response")
		}
		return nil, fmt.Errorf("no data field found in inventory service response: %+v", response)
	}

	// Now responseData could be the PaginationResponse or an array
	if dataMap, ok := responseData.(map[string]interface{}); ok {
		// The inventory service returns a PaginationResponse where the actual items are in dataMap["data"]
		inventoryData, ok := dataMap["data"].([]interface{})
		if ok {
			return inventoryData, nil
		}
		// Try alternative field name "Data" (uppercase)
		if inventoryData, ok := dataMap["Data"].([]interface{}); ok {
			return inventoryData, nil
		}
		return nil, fmt.Errorf("invalid inventory data format from inventory service")
	}

	// If responseData is directly an array
	if inventoryData, ok := responseData.([]interface{}); ok {
		return inventoryData, nil
	}

	return nil, fmt.Errorf("invalid response format from inventory service")
}