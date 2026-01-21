package utils

import (
	"strings"
	"unicode"
)

// Truncate truncates string to specified length
func Truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}

// CapitalizeFirst capitalizes first letter
func CapitalizeFirst(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

// TitleCase converts string to title case
func TitleCase(s string) string {
	return strings.Title(strings.ToLower(s))
}

// RemoveSpaces removes all spaces from string
func RemoveSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

// CountWords counts words in string
func CountWords(s string) int {
	return len(strings.Fields(s))
}

// IsAlphanumeric checks if string contains only letters and numbers
func IsAlphanumeric(s string) bool {
	for _, char := range s {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
			return false
		}
	}
	return true
}

// ContainsAny checks if string contains any of the substrings
func ContainsAny(s string, substrs []string) bool {
	for _, substr := range substrs {
		if strings.Contains(s, substr) {
			return true
		}
	}
	return false
}

// SanitizeString removes dangerous characters for SQL/XSS
func SanitizeString(s string) string {
	// Remove potentially dangerous characters
	dangerous := []string{"<", ">", "'", "\"", ";", "--", "/*", "*/"}
	for _, char := range dangerous {
		s = strings.ReplaceAll(s, char, "")
	}
	return strings.TrimSpace(s)
}