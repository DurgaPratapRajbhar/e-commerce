package utils

import (
	"time"
)

// APIResponse standard response structure
type APIResponse struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message,omitempty"`      // ✅ Added
	Data      interface{} `json:"data,omitempty"`
	Error     *ErrorInfo  `json:"error,omitempty"`
	Meta      *Meta       `json:"meta,omitempty"`         // ✅ Added
	Timestamp string      `json:"timestamp"`
	RequestID string      `json:"request_id"`
}

// ErrorInfo error details structure
type ErrorInfo struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

// Meta additional metadata
type Meta struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	Filters    interface{} `json:"filters,omitempty"`
	Sort       string      `json:"sort,omitempty"`
}

// SuccessResponse creates success response with message
func SuccessResponse(data interface{}, message string, requestID string) APIResponse {
	return APIResponse{
		Success:   true,
		Message:   message,    // ✅ Fixed - now using the parameter
		Data:      data,       // ✅ Fixed - now including data
		Timestamp: GetUTCNow(),
		RequestID: requestID,
	}
}

// DataResponse creates response with data (no message)
func DataResponse(data interface{}, requestID string) APIResponse {
	return APIResponse{
		Success:   true,
		Data:      data,       // ✅ Fixed
		Timestamp: GetUTCNow(),
		RequestID: requestID,
	}
}

// ErrorResponse creates error response
func ErrorResponse(code string, message string, details interface{}, requestID string) APIResponse {
	return APIResponse{
		Success: false,
		Error: &ErrorInfo{
			Code:    code,
			Message: message,
			Details: details,
		},
		Timestamp: GetUTCNow(),
		RequestID: requestID,
	}
}

// ValidationErrorResponse creates validation error response
func ValidationErrorResponse(fieldErrors []ValidationError, requestID string) APIResponse {
	return APIResponse{
		Success: false,
		Error: &ErrorInfo{
			Code:    ErrValidationFailed,
			Message: "Validation failed",
			Details: map[string]interface{}{
				"fields": fieldErrors,
			},
		},
		Timestamp: GetUTCNow(),
		RequestID: requestID,
	}
}

// PaginatedResponse creates paginated response
func PaginatedResponse(items interface{}, pagination *Pagination, message string, requestID string) APIResponse {
	return APIResponse{
		Success: true,
		Message: message,
		Data: map[string]interface{}{
			"items": items,
		},
		Meta: &Meta{              // ✅ Fixed
			Pagination: pagination,
		},
		Timestamp: GetUTCNow(),
		RequestID: requestID,
	}
}

// GetUTCNow returns current UTC timestamp in ISO 8601 format
func GetUTCNow() string {
	return time.Now().UTC().Format(time.RFC3339)
}