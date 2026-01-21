package utils

import (
	"fmt"
	"strings"
)

// NormalizeEmail normalizes email address
func NormalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}

// MaskEmail masks email for privacy (u***@example.com)
func MaskEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return email
	}

	username := parts[0]
	domain := parts[1]

	if len(username) <= 2 {
		return fmt.Sprintf("%s***@%s", string(username[0]), domain)
	}

	return fmt.Sprintf("%s***@%s", string(username[0]), domain)
}

// GetEmailDomain extracts domain from email
func GetEmailDomain(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
}

// IsDisposableEmail checks if email is from disposable email service
func IsDisposableEmail(email string) bool {
	disposableDomains := []string{
		"tempmail.com", "guerrillamail.com", "10minutemail.com",
		"throwaway.email", "mailinator.com",
	}

	domain := GetEmailDomain(email)
	for _, disposable := range disposableDomains {
		if domain == disposable {
			return true
		}
	}
	return false
}