package utils

import (
	"errors"
	"fmt"
)

// Error codes
const (
	// Authentication errors (1xxx)
	ErrUnauthorized        = "AUTH_1001"
	ErrInvalidToken        = "AUTH_1002"
	ErrTokenExpired        = "AUTH_1003"
	ErrInvalidCredentials  = "AUTH_1004"
	ErrEmailNotVerified    = "AUTH_1005"
	ErrForbidden           = "AUTH_1006"

	// Validation errors (2xxx)
	ErrValidationFailed    = "VAL_2001"
	ErrInvalidInput        = "VAL_2002"
	ErrMissingField        = "VAL_2003"
	ErrInvalidFormat       = "VAL_2004"

	// Resource errors (3xxx)
	ErrNotFound            = "RES_3001"
	ErrAlreadyExists       = "RES_3002"
	ErrConflict            = "RES_3003"

	// Business logic errors (4xxx)
	ErrInsufficientStock   = "BIZ_4001"
	ErrInvalidOrder        = "BIZ_4002"
	ErrPaymentFailed       = "BIZ_4003"
	ErrInvalidDiscount     = "BIZ_4004"

	// Rate limit errors (6xxx)
	ErrRateLimitExceeded   = "RL_6001"
	
	// Server errors (5xxx)
	ErrInternalServer      = "SRV_5001"
	ErrDatabaseError       = "SRV_5002"
	ErrExternalService     = "SRV_5003"
)

// Error messages
var ErrorMessages = map[string]string{
	ErrUnauthorized:       "Unauthorized access",
	ErrInvalidToken:       "Invalid or malformed token",
	ErrTokenExpired:       "Token has expired",
	ErrInvalidCredentials: "Invalid email or password",
	ErrEmailNotVerified:   "Email not verified",
	ErrForbidden:          "Forbidden access",
	
	ErrValidationFailed:   "Validation failed",
	ErrInvalidInput:       "Invalid input provided",
	ErrMissingField:       "Required field is missing",
	ErrInvalidFormat:      "Invalid format",
	
	ErrNotFound:           "Resource not found",
	ErrAlreadyExists:      "Resource already exists",
	ErrConflict:           "Resource conflict",
	
	ErrInsufficientStock:  "Insufficient stock available",
	ErrInvalidOrder:       "Invalid order",
	ErrPaymentFailed:      "Payment processing failed",
	ErrInvalidDiscount:    "Invalid discount code",
	
	ErrRateLimitExceeded:  "Rate limit exceeded",
	
	ErrInternalServer:     "Internal server error",
	ErrDatabaseError:      "Database operation failed",
	ErrExternalService:    "External service error",
}

// GetErrorMessage returns message for error code
func GetErrorMessage(code string) string {
	if msg, exists := ErrorMessages[code]; exists {
		return msg
	}
	return "Unknown error"
}

// GetError returns a Go error for an error code
func GetError(code string) error {
	return errors.New(GetErrorMessage(code))
}

// GetErrorf returns a formatted Go error for an error code
func GetErrorf(code string, args ...interface{}) error {
	return fmt.Errorf(GetErrorMessage(code), args...)
}

// ValidationError represents validation error details
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidationErrors multiple validation errors
type ValidationErrors struct {
	Errors []ValidationError `json:"errors"`
}