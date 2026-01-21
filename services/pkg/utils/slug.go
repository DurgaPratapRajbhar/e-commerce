package utils

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// GenerateSlug generates URL-friendly slug from string
func GenerateSlug(text string) string {
	// Convert to lowercase
	slug := strings.ToLower(text)

	// Replace spaces with hyphens
	slug = strings.ReplaceAll(slug, " ", "-")

	// Remove special characters
	reg := regexp.MustCompile(`[^a-z0-9-]`)
	slug = reg.ReplaceAllString(slug, "")

	// Remove multiple consecutive hyphens
	reg = regexp.MustCompile(`-+`)
	slug = reg.ReplaceAllString(slug, "-")

	// Trim hyphens from start and end
	slug = strings.Trim(slug, "-")

	return slug
}

// GenerateUniqueSlug generates unique slug with timestamp
func GenerateUniqueSlug(text string) string {
	baseSlug := GenerateSlug(text)
	timestamp := time.Now().Unix()
	return fmt.Sprintf("%s-%d", baseSlug, timestamp)
}

// GenerateSKU generates product SKU
func GenerateSKU(prefix string, id int) string {
	return fmt.Sprintf("%s-%06d", strings.ToUpper(prefix), id)
}

// GenerateOrderNumber generates order number
func GenerateOrderNumber() string {
	now := time.Now()
	return fmt.Sprintf("ORD-%s-%d", now.Format("20060102"), now.Unix()%100000)
}