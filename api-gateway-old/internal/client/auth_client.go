package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/DurgaPratapRajbhar/e-commerce/pkg/config"
)

// AuthClient handles communication with the auth service
type AuthClient struct {
	baseURL string
	timeout time.Duration
}

// NewAuthClient creates a new auth service client
func NewAuthClient(cfg *config.Config) *AuthClient {
	return &AuthClient{
		baseURL: cfg.Services.AuthService.URL,
		timeout: cfg.Services.AuthService.Timeout,
	}
}

// makeRequest is a generic method to make HTTP requests
func (c *AuthClient) makeRequest(ctx context.Context, method, endpoint string, authToken string, body interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s%s", c.baseURL, endpoint)

	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if authToken != "" {
		req.Header.Set("Authorization", authToken)
	}

	client := &http.Client{Timeout: c.timeout}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Check HTTP status code
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(respBody))
	}

	var response map[string]interface{}
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Check if it's an error response
	if success, ok := response["success"].(bool); ok && !success {
		return nil, fmt.Errorf("service error: %v", response["message"])
	}

	return response, nil
}

// extractData extracts data from response with multiple fallback strategies
func (c *AuthClient) extractData(response map[string]interface{}) (interface{}, error) {
	// Strategy 1: response["data"]["data"]
	if data, ok := response["data"].(map[string]interface{}); ok {
		if innerData, ok := data["data"]; ok {
			return innerData, nil
		}
		if innerData, ok := data["Data"]; ok {
			return innerData, nil
		}
		return data, nil
	}

	// Strategy 2: response["Data"]["Data"]
	if data, ok := response["Data"].(map[string]interface{}); ok {
		if innerData, ok := data["Data"]; ok {
			return innerData, nil
		}
		if innerData, ok := data["data"]; ok {
			return innerData, nil
		}
		return data, nil
	}

	// Strategy 3: response["data"] is directly array
	if data, ok := response["data"]; ok {
		return data, nil
	}

	return nil, fmt.Errorf("no valid data field found in response")
}

// GetUsers fetches users with pagination
func (c *AuthClient) GetUsers(ctx context.Context, authToken string, page, limit string) ([]interface{}, error) {
	endpoint := fmt.Sprintf("/auth/admin/users?page=%s&limit=%s", page, limit)
	
	response, err := c.makeRequest(ctx, "GET", endpoint, authToken, nil)
	if err != nil {
		return nil, err
	}

	data, err := c.extractData(response)
	if err != nil {
		return nil, err
	}

	// Convert to array
	if arr, ok := data.([]interface{}); ok {
		return arr, nil
	}

	return nil, fmt.Errorf("invalid data format: expected array")
}

// GetUserByID fetches a single user by ID
func (c *AuthClient) GetUserByID(ctx context.Context, authToken, userID string) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("/auth/admin/users/%s", userID)
	
	response, err := c.makeRequest(ctx, "GET", endpoint, authToken, nil)
	if err != nil {
		return nil, err
	}

	data, err := c.extractData(response)
	if err != nil {
		return nil, err
	}

	if userMap, ok := data.(map[string]interface{}); ok {
		return userMap, nil
	}

	return nil, fmt.Errorf("invalid data format: expected object")
}

// CreateUser creates a new user
func (c *AuthClient) CreateUser(ctx context.Context, authToken string, userData map[string]interface{}) (map[string]interface{}, error) {
	endpoint := "/auth/admin/users"
	
	response, err := c.makeRequest(ctx, "POST", endpoint, authToken, userData)
	if err != nil {
		return nil, err
	}

	data, err := c.extractData(response)
	if err != nil {
		return nil, err
	}

	if userMap, ok := data.(map[string]interface{}); ok {
		return userMap, nil
	}

	return nil, fmt.Errorf("invalid data format: expected object")
}

// UpdateUser updates an existing user
func (c *AuthClient) UpdateUser(ctx context.Context, authToken, userID string, userData map[string]interface{}) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("/auth/admin/users/%s", userID)
	
	response, err := c.makeRequest(ctx, "PUT", endpoint, authToken, userData)
	if err != nil {
		return nil, err
	}

	data, err := c.extractData(response)
	if err != nil {
		return nil, err
	}

	if userMap, ok := data.(map[string]interface{}); ok {
		return userMap, nil
	}

	return nil, fmt.Errorf("invalid data format: expected object")
}

// DeleteUser deletes a user
func (c *AuthClient) DeleteUser(ctx context.Context, authToken, userID string) error {
	endpoint := fmt.Sprintf("/auth/admin/users/%s", userID)
	
	_, err := c.makeRequest(ctx, "DELETE", endpoint, authToken, nil)
	return err
}

// UpdateUserStatus updates user status
func (c *AuthClient) UpdateUserStatus(ctx context.Context, authToken, userID string, status string) error {
	endpoint := fmt.Sprintf("/auth/admin/users/%s/status", userID)
	body := map[string]interface{}{"status": status}
	
	_, err := c.makeRequest(ctx, "PATCH", endpoint, authToken, body)
	return err
}