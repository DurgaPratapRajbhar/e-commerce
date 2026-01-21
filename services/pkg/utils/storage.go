package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ImageBaseDir defines the base directory for storing images
const ImageBaseDir = "github.com/DurgaPratapRajbhar/e-commerce/image/"

// AllowedImageFormats defines supported image extensions
var AllowedImageFormats = []string{".jpg", ".jpeg", ".png", ".gif"}

// EnsureImageDirExists checks if the image directory exists, if not, creates it
func EnsureImageDirExists() error {
	if err := os.MkdirAll(ImageBaseDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	return nil
}

// IsValidImageFormat checks if the file extension is allowed
func IsValidImageFormat(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	for _, allowedExt := range AllowedImageFormats {
		if ext == allowedExt {
			return true
		}
	}
	return false
}

// StoreImage saves an image file to `ImageBaseDir`
func StoreImage(file []byte, filename string) (string, error) {
	// Ensure directory exists
	if err := EnsureImageDirExists(); err != nil {
		return "", err
	}

	// Validate image format
	if !IsValidImageFormat(filename) {
		return "", fmt.Errorf("invalid image format: %s", filepath.Ext(filename))
	}

	// Define destination path
	destPath := filepath.Join(ImageBaseDir, filename)

	// Write file
	if err := os.WriteFile(destPath, file, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to save file: %w", err)
	}

	return destPath, nil
}

// DeleteImage removes an image file from the storage directory
func DeleteImage(imagePath string) error {
	if err := os.Remove(imagePath); err != nil {
		return fmt.Errorf("failed to delete file %s: %w", imagePath, err)
	}
	return nil
}
