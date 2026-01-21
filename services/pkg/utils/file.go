package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

// FileValidator validates file uploads
type FileValidator struct {
	MaxFileSize  int64
	AllowedTypes []string
}

// NewFileValidator creates a new file validator
func NewFileValidator(maxFileSize int64, allowedTypes []string) *FileValidator {
	return &FileValidator{
		MaxFileSize:  maxFileSize,
		AllowedTypes: allowedTypes,
	}
}

// ValidateFile validates a file upload
func (fv *FileValidator) ValidateFile(fileHeader *multipart.FileHeader) error {
	// Check file size
	if fileHeader.Size > fv.MaxFileSize {
		return fmt.Errorf("file size exceeds maximum allowed size of %d bytes", fv.MaxFileSize)
	}

	// Check file type
	fileType := fileHeader.Header.Get("Content-Type")
	isAllowed := false
	for _, allowedType := range fv.AllowedTypes {
		if strings.HasPrefix(fileType, allowedType) {
			isAllowed = true
			break
		}
	}

	if !isAllowed {
		return fmt.Errorf("file type '%s' is not allowed", fileType)
	}

	return nil
}

// SaveFile saves an uploaded file to the specified path
func SaveFile(fileHeader *multipart.FileHeader, destPath string) error {
	// Create destination directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	// Open the uploaded file
	src, err := fileHeader.Open()
	if err != nil {
		return fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer src.Close()

	// Create the destination file
	dest, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dest.Close()

	// Copy the file contents
	if _, err := io.Copy(dest, src); err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}

	return nil
}

// FileExists checks if a file exists
func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

// DeleteFile deletes a file
func DeleteFile(filePath string) error {
	return os.Remove(filePath)
}

// GetFileSize returns the size of a file
func GetFileSize(filePath string) (int64, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

// GetFileExtension returns the file extension without the dot
func GetFileExtension(filename string) string {
	return strings.TrimPrefix(filepath.Ext(filename), ".")
}

// GetFileNameWithoutExtension returns the filename without extension
func GetFileNameWithoutExtension(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

// SanitizeFileName removes potentially dangerous characters from a filename
func SanitizeFileName(filename string) string {
	// Remove path separators to prevent directory traversal
	filename = strings.ReplaceAll(filename, "/", "")
	filename = strings.ReplaceAll(filename, "\\", "")
	
	// Remove other potentially dangerous characters
	filename = strings.ReplaceAll(filename, "..", "")
	
	return filename
}