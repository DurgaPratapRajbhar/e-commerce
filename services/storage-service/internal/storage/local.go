package storage

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type LocalStorage struct {
	BasePath string
	BaseURL  string
}

func NewLocalStorage(basePath, baseURL string) *LocalStorage {
	return &LocalStorage{
		BasePath: basePath,
		BaseURL:  baseURL,
	}
}

func (s *LocalStorage) SaveFile(file *multipart.FileHeader, category, subcategory string) (string, error) {
	// Create the full path for the file
	uploadPath := filepath.Join(s.BasePath, category, subcategory)

	// Create directory if it doesn't exist
	if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create upload directory: %w", err)
	}

	// Generate unique filename
	filename := file.Filename
	uniqueFilename := generateUniqueFilename(filename)
	fullPath := filepath.Join(uploadPath, uniqueFilename)

	// Save the file
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer src.Close()

	dst, err := os.Create(fullPath)
	if err != nil {
		return "", fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("failed to save file: %w", err)
	}

	return filepath.Join(category, subcategory, uniqueFilename), nil
}

func (s *LocalStorage) GetFile(filepath string) ([]byte, error) {
	fullPath := filepath.Join(s.BasePath, filepath)
	data, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	return data, nil
}

func (s *LocalStorage) DeleteFile(filepath string) error {
	fullPath := filepath.Join(s.BasePath, filepath)
	if err := os.Remove(fullPath); err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}
	return nil
}

func (s *LocalStorage) FileExists(filepath string) bool {
	fullPath := filepath.Join(s.BasePath, filepath)
	_, err := os.Stat(fullPath)
	return err == nil
}

// generateUniqueFilename generates a unique filename by adding timestamp
func generateUniqueFilename(originalName string) string {
	// Get file extension and name
	ext := filepath.Ext(originalName)
	name := strings.TrimSuffix(originalName, ext)
	
	// Generate unique filename with timestamp
	timestamp := fmt.Sprintf("%d_%s%s", getCurrentTimestamp(), name, ext)
	return timestamp
}

// getCurrentTimestamp returns current timestamp as int64
func getCurrentTimestamp() int64 {
	return time.Now().Unix()
}