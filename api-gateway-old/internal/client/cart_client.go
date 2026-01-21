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

// CartClient handles communication with the cart service
type CartClient struct {
	baseURL string
	timeout time.Duration
}

// NewCartClient creates a new cart service client
func NewCartClient(cfg *config.Config) *CartClient {
	return &CartClient{
		baseURL: cfg.Services.CartService.URL,
		timeout: cfg.Services.CartService.Timeout,
	}
}

// GetCartItems fetches cart items from cart service
func (c *CartClient) GetCartItems(ctx context.Context, authToken string) ([]interface{}, error) {
	url := fmt.Sprintf("%s/cart/items", c.baseURL)

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
		return nil, fmt.Errorf("cart service returned error: %v", response["message"])
	}

	// Try to get the data field
	responseData, ok := response["data"]
	if !ok {
		// Check for alternative field names
		if dataAlt, okAlt := response["Data"].(map[string]interface{}); okAlt {
			// If data is directly an array
			if cartData, okArray := dataAlt["items"].([]interface{}); okArray {
				return cartData, nil
			}
			// If data has "Data" field
			if cartData, okArray := dataAlt["Data"].([]interface{}); okArray {
				return cartData, nil
			}
			// If data has "data" field
			if cartData, okArray := dataAlt["data"].([]interface{}); okArray {
				return cartData, nil
			}
			return nil, fmt.Errorf("invalid data format in cart service response")
		}
		return nil, fmt.Errorf("no data field found in cart service response: %+v", response)
	}

	// Now responseData could be the PaginationResponse or an array
	if dataMap, ok := responseData.(map[string]interface{}); ok {
		// The cart service returns data where the actual items are in dataMap["data"]
		cartData, ok := dataMap["data"].([]interface{})
		if ok {
			return cartData, nil
		}
		// Try alternative field name "Data" (uppercase)
		if cartData, ok := dataMap["Data"].([]interface{}); ok {
			return cartData, nil
		}
		return nil, fmt.Errorf("invalid cart data format from cart service")
	}

	// If responseData is directly an array
	if cartData, ok := responseData.([]interface{}); ok {
		return cartData, nil
	}

	return nil, fmt.Errorf("invalid response format from cart service")
}