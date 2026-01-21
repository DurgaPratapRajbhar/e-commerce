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

// ShippingClient handles communication with the shipping service
type ShippingClient struct {
	baseURL string
	timeout time.Duration
}

// NewShippingClient creates a new shipping service client
func NewShippingClient(cfg *config.Config) *ShippingClient {
	return &ShippingClient{
		baseURL: cfg.Services.ShippingService.URL,
		timeout: cfg.Services.ShippingService.Timeout,
	}
}

// GetShipments fetches shipments with pagination from shipping service
func (c *ShippingClient) GetShipments(ctx context.Context, authToken string, page, limit string) ([]interface{}, error) {
	url := fmt.Sprintf("%s/shipment/shipments?page=%s&limit=%s", c.baseURL, page, limit)

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
		return nil, fmt.Errorf("shipping service returned error: %v", response["message"])
	}

	// Try to get the data field
	responseData, ok := response["data"]
	if !ok {
		// Check for alternative field names
		if dataAlt, okAlt := response["Data"].(map[string]interface{}); okAlt {
			// If data is directly an array
			if shipmentsData, okArray := dataAlt["shipments"].([]interface{}); okArray {
				return shipmentsData, nil
			}
			// If data has "Data" field
			if shipmentsData, okArray := dataAlt["Data"].([]interface{}); okArray {
				return shipmentsData, nil
			}
			// If data has "data" field
			if shipmentsData, okArray := dataAlt["data"].([]interface{}); okArray {
				return shipmentsData, nil
			}
			return nil, fmt.Errorf("invalid data format in shipping service response")
		}
		return nil, fmt.Errorf("no data field found in shipping service response: %+v", response)
	}

	// Now responseData could be the PaginationResponse or an array
	if dataMap, ok := responseData.(map[string]interface{}); ok {
		// The shipping service returns a PaginationResponse where the actual shipments are in dataMap["data"]
		shipmentsData, ok := dataMap["data"].([]interface{})
		if ok {
			return shipmentsData, nil
		}
		// Try alternative field name "Data" (uppercase)
		if shipmentsData, ok := dataMap["Data"].([]interface{}); ok {
			return shipmentsData, nil
		}
		return nil, fmt.Errorf("invalid shipments data format from shipping service")
	}

	// If responseData is directly an array
	if shipmentsData, ok := responseData.([]interface{}); ok {
		return shipmentsData, nil
	}

	return nil, fmt.Errorf("invalid response format from shipping service")
}