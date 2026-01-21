package utils

import (
	"fmt"
	"regexp"
	"strings"
)

// FormatIndianPhone formats Indian phone number
func FormatIndianPhone(phone string) string {
	// Remove all non-digit characters
	phone = regexp.MustCompile(`\D`).ReplaceAllString(phone, "")

	// Remove country code if present
	if strings.HasPrefix(phone, "91") && len(phone) == 12 {
		phone = phone[2:]
	}

	if len(phone) == 10 {
		return fmt.Sprintf("+91 %s %s %s", 
			phone[0:5], phone[5:8], phone[8:10])
	}

	return phone
}

// ValidateIndianPhone validates Indian phone number
func ValidateIndianPhone(phone string) bool {
	phone = regexp.MustCompile(`\D`).ReplaceAllString(phone, "")
	
	// Remove +91 if present
	if strings.HasPrefix(phone, "91") {
		phone = phone[2:]
	}

	// Must be 10 digits starting with 6-9
	phoneRegex := regexp.MustCompile(`^[6-9]\d{9}$`)
	return phoneRegex.MatchString(phone)
}

// MaskPhone masks phone number (98***45678)
func MaskPhone(phone string) string {
	if len(phone) < 10 {
		return phone
	}

	return fmt.Sprintf("%s***%s", 
		phone[:2], phone[len(phone)-5:])
}