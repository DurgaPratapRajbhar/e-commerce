package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// EnsureDir ensures that the directory exists, creating it if necessary
func EnsureDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

// SanitizePath sanitizes a file path to prevent directory traversal
func SanitizePath(path string) string {
	// Clean the path to remove any relative path components
	cleanPath := filepath.Clean(path)
	
	// Remove any leading separators to prevent absolute paths
	cleanPath = strings.TrimPrefix(cleanPath, string(filepath.Separator))
	
	return cleanPath
}

// IsValidPath checks if a path is valid (doesn't contain dangerous patterns)
func IsValidPath(path string) bool {
	// Check for directory traversal attempts
	if strings.Contains(path, "../") || strings.Contains(path, "..\\") {
		return false
	}
	
	// Check if path is clean (no relative components)
	cleanPath := filepath.Clean(path)
	if cleanPath != path {
		return false
	}
	
	return true
}

// JoinPath joins path elements safely
func JoinPath(elem ...string) string {
	if len(elem) == 0 {
		return ""
	}
	
	result := elem[0]
	for _, e := range elem[1:] {
		result = filepath.Join(result, e)
	}
	
	return result
}

// GetFileExtension returns the file extension without the dot
func GetFileExtension(filename string) string {
	ext := filepath.Ext(filename)
	if ext != "" && len(ext) > 1 {
		return ext[1:] // Remove the leading dot
	}
	return ""
}

// GetFileNameWithoutExtension returns the filename without extension
func GetFileNameWithoutExtension(filename string) string {
	ext := filepath.Ext(filename)
	return strings.TrimSuffix(filename, ext)
}