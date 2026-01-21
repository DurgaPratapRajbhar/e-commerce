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

// ProductClient handles communication with the product service
type ProductClient struct {
	baseURL string
	timeout time.Duration
}

// NewProductClient creates a new product service client
func NewProductClient(cfg *config.Config) *ProductClient {
	return &ProductClient{
		baseURL: cfg.Services.ProductService.URL,
		timeout: cfg.Services.ProductService.Timeout,
	}
}

// GetProducts fetches products with pagination from product service
func (c *ProductClient) GetProducts(ctx context.Context, authToken string, page, limit string) ([]interface{}, error) {
	url := fmt.Sprintf("%s/product/products?page=%s&limit=%s", c.baseURL, page, limit)

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
		return nil, fmt.Errorf("product service returned error: %v", response["message"])
	}

	// Try to get the data field
	responseData, ok := response["data"]
	if !ok {
		// Check for alternative field names
		if dataAlt, okAlt := response["Data"].(map[string]interface{}); okAlt {
			// If data is directly an array
			if productsData, okArray := dataAlt["products"].([]interface{}); okArray {
				return productsData, nil
			}
			// If data has "Data" field
			if productsData, okArray := dataAlt["Data"].([]interface{}); okArray {
				return productsData, nil
			}
			// If data has "data" field
			if productsData, okArray := dataAlt["data"].([]interface{}); okArray {
				return productsData, nil
			}
			return nil, fmt.Errorf("invalid data format in product service response")
		}
		return nil, fmt.Errorf("no data field found in product service response: %+v", response)
	}

	// Now responseData could be the PaginationResponse or an array
	if dataMap, ok := responseData.(map[string]interface{}); ok {
		// The product service returns a PaginationResponse where the actual products are in dataMap["data"]
		productsData, ok := dataMap["data"].([]interface{})
		if ok {
			return productsData, nil
		}
		// Try alternative field name "Data" (uppercase)
		if productsData, ok := dataMap["Data"].([]interface{}); ok {
			return productsData, nil
		}
		return nil, fmt.Errorf("invalid products data format from product service")
	}

	// If responseData is directly an array
	if productsData, ok := responseData.([]interface{}); ok {
		return productsData, nil
	}

	return nil, fmt.Errorf("invalid response format from product service")
}

// GetProductsByIDs fetches products by their IDs from product service
func (c *ProductClient) GetProductsByIDs(ctx context.Context, authToken string, productIDs []string) ([]interface{}, error) {
	// Build query string with all product IDs
	url := fmt.Sprintf("%s/product/products-by-ids?ids=%s", c.baseURL, "")
	
	// For now, we'll make a simple GET request with comma-separated IDs
	// In a real implementation, you might want to use POST for large ID lists
	idsStr := ""
	for i, id := range productIDs {
		if i == 0 {
			idsStr = id
		} else {
			idsStr = fmt.Sprintf("%s,%s", idsStr, id)
		}
	}
	url = fmt.Sprintf("%s/product/products-by-ids?ids=%s", c.baseURL, idsStr)

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
		return nil, fmt.Errorf("product service returned error: %v", response["message"])
	}

	// Try to get the data field
	responseData, ok := response["data"]
	if !ok {
		return nil, fmt.Errorf("no data field found in product service response: %+v", response)
	}

	// If responseData is an array of products
	if productsData, ok := responseData.([]interface{}); ok {
		return productsData, nil
	}

	// If responseData is a map, try to extract products
	if dataMap, ok := responseData.(map[string]interface{}); ok {
		if productsData, ok := dataMap["data"].([]interface{}); ok {
			return productsData, nil
		}
		// Try alternative field name "Data"
		if productsData, ok := dataMap["Data"].([]interface{}); ok {
			return productsData, nil
		}
	}

	return nil, fmt.Errorf("invalid response format from product service")
}