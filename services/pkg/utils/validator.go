package utils

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

// ValidateEmail validates an email address
func ValidateEmail(email string) bool {
	if len(email) < 3 || len(email) > 254 {
		return false
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// ValidatePassword validates a password based on security requirements
func ValidatePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}

	if len(password) > 128 {
		return fmt.Errorf("password must not exceed 128 characters")
	}

	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	if len(password) >= 8 {
		hasMinLen = true
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasMinLen {
		return fmt.Errorf("password must be at least 8 characters long")
	}
	if !hasUpper {
		return fmt.Errorf("password must contain at least one uppercase letter")
	}
	if !hasLower {
		return fmt.Errorf("password must contain at least one lowercase letter")
	}
	if !hasNumber {
		return fmt.Errorf("password must contain at least one number")
	}
	if !hasSpecial {
		return fmt.Errorf("password must contain at least one special character")
	}

	return nil
}

// ValidatePhone validates a phone number
func ValidatePhone(phone string) bool {
	// Remove all non-digit characters
	re := regexp.MustCompile(`[^\d+]`)
	cleaned := re.ReplaceAllString(phone, "")

	// Check if it has between 10 and 15 digits (international format)
	if len(cleaned) < 10 || len(cleaned) > 15 {
		return false
	}

	// Basic validation for phone number format
	phoneRegex := regexp.MustCompile(`^[\+]?[1-9][\d]{0,15}$`)
	return phoneRegex.MatchString(cleaned)
}

// ValidateRequired checks if a string is not empty after trimming
func ValidateRequired(value string) bool {
	return strings.TrimSpace(value) != ""
}

// ValidateMinLength checks if a string meets minimum length requirement
func ValidateMinLength(value string, minLength int) bool {
	return len(strings.TrimSpace(value)) >= minLength
}

// ValidateMaxLength checks if a string meets maximum length requirement
func ValidateMaxLength(value string, maxLength int) bool {
	return len(strings.TrimSpace(value)) <= maxLength
}

// ValidateBetweenLength checks if a string length is between min and max
func ValidateBetweenLength(value string, minLength, maxLength int) bool {
	length := len(strings.TrimSpace(value))
	return length >= minLength && length <= maxLength
}

// ValidateURL validates if a string is a valid URL
func ValidateURL(url string) bool {
	urlRegex := regexp.MustCompile(`^(https?://)?([a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}(/.*)?$`)
	return urlRegex.MatchString(url)
}

// ValidateAlphaNumeric validates if a string contains only alphanumeric characters
func ValidateAlphaNumeric(value string) bool {
	alphaNumRegex := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	return alphaNumRegex.MatchString(value)
}

// ValidateAlphaNumericSpace validates if a string contains only alphanumeric characters and spaces
func ValidateAlphaNumericSpace(value string) bool {
	alphaNumSpaceRegex := regexp.MustCompile(`^[a-zA-Z0-9 ]+$`)
	return alphaNumSpaceRegex.MatchString(value)
}

// ValidateAlphaNumericHyphenUnderscore validates if a string contains alphanumeric, hyphens, and underscores
func ValidateAlphaNumericHyphenUnderscore(value string) bool {
	alphaNumHyphenUnderscoreRegex := regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
	return alphaNumHyphenUnderscoreRegex.MatchString(value)
}

// ValidateStringInSlice checks if a string is present in a slice
func ValidateStringInSlice(value string, validValues []string) bool {
	for _, validValue := range validValues {
		if value == validValue {
			return true
		}
	}
	return false
}

// ParseValidationErrors parses gin validation error messages and extracts field names and messages
func ParseValidationErrors(errorMsg string) []ValidationError {
	var validationErrors []ValidationError
	
	// Regex to match the pattern: Key: 'StructName.FieldName' Error:Field validation for 'FieldName' failed on the 'tag' tag
	re := regexp.MustCompile(`Key: '([^']+)\.([^']+)' Error:Field validation for '([^']+)' failed on the '([^']*)' tag`)
	matches := re.FindAllStringSubmatch(errorMsg, -1)
	
	for _, match := range matches {
		if len(match) >= 4 {
			fieldName := strings.ToLower(match[2]) // Convert to lowercase for consistency
			validationTag := match[4]
			
			// Create a user-friendly message based on the validation tag
			message := createValidationMessage(fieldName, validationTag)
			
			validationErrors = append(validationErrors, ValidationError{
				Field:   fieldName,
				Message: message,
			})
		}
	}
	
	// If no matches were found, use the original error message
	if len(validationErrors) == 0 {
		validationErrors = append(validationErrors, ValidationError{
			Field:   "request_body",
			Message: errorMsg,
		})
	}
	
	return validationErrors
}

// createValidationMessage creates a user-friendly validation message based on the validation tag
func createValidationMessage(fieldName, validationTag string) string {
	switch validationTag {
	case "required":
		return fieldName + " is required"
	case "email":
		return fieldName + " must be a valid email"
	case "min":
		return fieldName + " is too short"
	case "max":
		return fieldName + " is too long"
	case "len":
		return fieldName + " has invalid length"
	default:
		return "Invalid " + fieldName
	}
}