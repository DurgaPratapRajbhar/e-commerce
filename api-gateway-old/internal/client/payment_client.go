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

// PaymentClient handles communication with the payment service
type PaymentClient struct {
	baseURL string
	timeout time.Duration
}

// NewPaymentClient creates a new payment service client
func NewPaymentClient(cfg *config.Config) *PaymentClient {
	return &PaymentClient{
		baseURL: cfg.Services.PaymentService.URL,
		timeout: cfg.Services.PaymentService.Timeout,
	}
}

// GetPayments fetches payments with pagination from payment service
func (c *PaymentClient) GetPayments(ctx context.Context, authToken string, page, limit string) ([]interface{}, error) {
	url := fmt.Sprintf("%s/payment/payments?page=%s&limit=%s", c.baseURL, page, limit)

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
		return nil, fmt.Errorf("payment service returned error: %v", response["message"])
	}

	// Try to get the data field
	responseData, ok := response["data"]
	if !ok {
		// Check for alternative field names
		if dataAlt, okAlt := response["Data"].(map[string]interface{}); okAlt {
			// If data is directly an array
			if paymentsData, okArray := dataAlt["payments"].([]interface{}); okArray {
				return paymentsData, nil
			}
			// If data has "Data" field
			if paymentsData, okArray := dataAlt["Data"].([]interface{}); okArray {
				return paymentsData, nil
			}
			// If data has "data" field
			if paymentsData, okArray := dataAlt["data"].([]interface{}); okArray {
				return paymentsData, nil
			}
			return nil, fmt.Errorf("invalid data format in payment service response")
		}
		return nil, fmt.Errorf("no data field found in payment service response: %+v", response)
	}

	// Now responseData could be the PaginationResponse or an array
	if dataMap, ok := responseData.(map[string]interface{}); ok {
		// The payment service returns a PaginationResponse where the actual payments are in dataMap["data"]
		paymentsData, ok := dataMap["data"].([]interface{})
		if ok {
			return paymentsData, nil
		}
		// Try alternative field name "Data" (uppercase)
		if paymentsData, ok := dataMap["Data"].([]interface{}); ok {
			return paymentsData, nil
		}
		return nil, fmt.Errorf("invalid payments data format from payment service")
	}

	// If responseData is directly an array
	if paymentsData, ok := responseData.([]interface{}); ok {
		return paymentsData, nil
	}

	return nil, fmt.Errorf("invalid response format from payment service")
}