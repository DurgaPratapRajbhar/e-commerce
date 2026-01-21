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

// OrderClient handles communication with the order service
type OrderClient struct {
	baseURL string
	timeout time.Duration
}

// NewOrderClient creates a new order service client
func NewOrderClient(cfg *config.Config) *OrderClient {
	return &OrderClient{
		baseURL: cfg.Services.OrderService.URL,
		timeout: cfg.Services.OrderService.Timeout,
	}
}

// GetOrders fetches orders with pagination from order service
func (c *OrderClient) GetOrders(ctx context.Context, authToken string, page, limit string) ([]interface{}, error) {
	url := fmt.Sprintf("%s/order/orders?page=%s&limit=%s", c.baseURL, page, limit)

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
		return nil, fmt.Errorf("order service returned error: %v", response["message"])
	}

	// Try to get the data field
	responseData, ok := response["data"]
	if !ok {
		// Check for alternative field names
		if dataAlt, okAlt := response["Data"].(map[string]interface{}); okAlt {
			// If data is directly an array
			if ordersData, okArray := dataAlt["orders"].([]interface{}); okArray {
				return ordersData, nil
			}
			// If data has "Data" field
			if ordersData, okArray := dataAlt["Data"].([]interface{}); okArray {
				return ordersData, nil
			}
			// If data has "data" field
			if ordersData, okArray := dataAlt["data"].([]interface{}); okArray {
				return ordersData, nil
			}
			return nil, fmt.Errorf("invalid data format in order service response")
		}
		return nil, fmt.Errorf("no data field found in order service response: %+v", response)
	}

	// Now responseData could be the PaginationResponse or an array
	if dataMap, ok := responseData.(map[string]interface{}); ok {
		// The order service returns a PaginationResponse where the actual orders are in dataMap["data"]
		ordersData, ok := dataMap["data"].([]interface{})
		if ok {
			return ordersData, nil
		}
		// Try alternative field name "Data" (uppercase)
		if ordersData, ok := dataMap["Data"].([]interface{}); ok {
			return ordersData, nil
		}
		return nil, fmt.Errorf("invalid orders data format from order service")
	}

	// If responseData is directly an array
	if ordersData, ok := responseData.([]interface{}); ok {
		return ordersData, nil
	}

	return nil, fmt.Errorf("invalid response format from order service")
}

// GetOrdersByUserID fetches orders by user ID from order service
func (c *OrderClient) GetOrdersByUserID(ctx context.Context, userID, authToken string, page, limit string) ([]interface{}, error) {
	url := fmt.Sprintf("%s/order/users/%s/orders?page=%s&limit=%s", c.baseURL, userID, page, limit)

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
		return nil, fmt.Errorf("order service returned error: %v", response["message"])
	}

	// Try to get the data field
	responseData, ok := response["data"]
	if !ok {
		// Check for alternative field names
		if dataAlt, okAlt := response["Data"].([]interface{}); okAlt {
			return dataAlt, nil
		}
		return nil, fmt.Errorf("no data field found in order service response: %+v", response)
	}

	// If responseData is an array
	if ordersData, ok := responseData.([]interface{}); ok {
		return ordersData, nil
	}

	// If responseData is a map, try to extract orders
	if dataMap, ok := responseData.(map[string]interface{}); ok {
		if ordersData, ok := dataMap["data"].([]interface{}); ok {
			return ordersData, nil
		}
		// Try alternative field name "Data"
		if ordersData, ok := dataMap["Data"].([]interface{}); ok {
			return ordersData, nil
		}
	}

	return nil, fmt.Errorf("invalid response format from order service")
}