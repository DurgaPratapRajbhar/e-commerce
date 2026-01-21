package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/DurgaPratapRajbhar/e-commerce/pkg/config"
)

// UserClient handles communication with the user service
type UserClient struct {
	baseURL string
	timeout time.Duration
}

// NewUserClient creates a new user service client
func NewUserClient(cfg *config.Config) *UserClient {
	return &UserClient{
		baseURL: cfg.Services.UserService.URL,
		timeout: cfg.Services.UserService.Timeout,
	}
}

// GetProfilesByUserIDs fetches user profiles by multiple user IDs
func (c *UserClient) GetProfilesByUserIDs(ctx context.Context, authToken string, userIDs []string) ([]interface{}, error) {
	userIDsStr := strings.Join(userIDs, ",")
	url := fmt.Sprintf("%s/user/profiles/bulk?user_ids=%s", c.baseURL, userIDsStr)

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
		return nil, fmt.Errorf("user service returned error: %v", response["message"])
	}

	// Try to get the data field
	responseData, ok := response["data"]
	if !ok {
		// Check for alternative field names
		if dataAlt, okAlt := response["Data"].(map[string]interface{}); okAlt {
			// If data is directly an array
			if profilesData, okArray := dataAlt["profiles"].([]interface{}); okArray {
				return profilesData, nil
			}
			// If data has "Data" field
			if profilesData, okArray := dataAlt["Data"].([]interface{}); okArray {
				return profilesData, nil
			}
			// If data has "data" field
			if profilesData, okArray := dataAlt["data"].([]interface{}); okArray {
				return profilesData, nil
			}
			return nil, fmt.Errorf("invalid data format in user service response")
		}
		return nil, fmt.Errorf("no data field found in user service response: %+v", response)
	}

	// Now responseData could be the PaginationResponse or an array
	if dataMap, ok := responseData.(map[string]interface{}); ok {
		// The user service returns data in a similar format to auth service
		profilesData, ok := dataMap["data"].([]interface{})
		if ok {
			return profilesData, nil
		}
		// Try alternative field name "Data" (uppercase)
		if profilesData, ok := dataMap["Data"].([]interface{}); ok {
			return profilesData, nil
		}
		return nil, fmt.Errorf("invalid profiles data format from user service")
	}

	// If responseData is directly an array
	if profilesData, ok := responseData.([]interface{}); ok {
		return profilesData, nil
	}

	return nil, fmt.Errorf("invalid response format from user service")
}

// GetUsers fetches users with pagination from user service
func (c *UserClient) GetUsers(ctx context.Context, authToken string, page, limit string) ([]interface{}, error) {
	url := fmt.Sprintf("%s/user/users?page=%s&limit=%s", c.baseURL, page, limit)

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
		return nil, fmt.Errorf("user service returned error: %v", response["message"])
	}

	// Try to get the data field
	responseData, ok := response["data"]
	if !ok {
		// Check for alternative field names
		if dataAlt, okAlt := response["Data"].([]interface{}); okAlt {
			return dataAlt, nil
		}
		return nil, fmt.Errorf("no data field found in user service response: %+v", response)
	}

	// If responseData is an array
	if usersData, ok := responseData.([]interface{}); ok {
		return usersData, nil
	}

	// If responseData is a map, try to extract users
	if dataMap, ok := responseData.(map[string]interface{}); ok {
		if usersData, ok := dataMap["data"].([]interface{}); ok {
			return usersData, nil
		}
		// Try alternative field name "Data"
		if usersData, ok := dataMap["Data"].([]interface{}); ok {
			return usersData, nil
		}
	}

	return nil, fmt.Errorf("invalid response format from user service")
}

// GetUserByID fetches a user by ID from user service
func (c *UserClient) GetUserByID(ctx context.Context, userID, authToken string) (interface{}, error) {
	url := fmt.Sprintf("%s/user/%s", c.baseURL, userID)

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
		return nil, fmt.Errorf("user service returned error: %v", response["message"])
	}

	// Try to get the data field
	responseData, ok := response["data"]
	if !ok {
		return nil, fmt.Errorf("no data field found in user service response: %+v", response)
	}

	// If responseData is a user object
	if userData, ok := responseData.(map[string]interface{}); ok {
		return userData, nil
	}

	return nil, fmt.Errorf("invalid response format from user service")
}