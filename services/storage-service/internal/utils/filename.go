package utils

import (
	"crypto/rand"
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

// GenerateUniqueFilename generates a unique filename by adding timestamp and random string
func GenerateUniqueFilename(originalName string) string {
	ext := filepath.Ext(originalName)
	name := strings.TrimSuffix(originalName, ext)
	
	// Generate timestamp and random suffix
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	randomSuffix := generateRandomString(6)
	
	// Combine to create unique filename
	uniqueName := fmt.Sprintf("%s_%s_%s%s", name, timestamp, randomSuffix, ext)
	return uniqueName
}

// generateRandomString generates a random string of specified length
func generateRandomString(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		// Fallback to timestamp if random generation fails
		return fmt.Sprintf("rand%d", time.Now().UnixNano())
	}
	
	// Convert to hex and truncate to desired length
	return fmt.Sprintf("%x", bytes)[:length]
}