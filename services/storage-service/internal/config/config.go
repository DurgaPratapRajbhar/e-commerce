package config

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv  string
	Server  ServerConfig
	Storage StorageConfig
}

type ServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type StorageConfig struct {
	UploadPath      string
	BaseURL         string
	MaxFileSize     int64
	AllowedTypes    []string
	ResizeWidth     int
	ResizeHeight    int
	ThumbnailWidth  int
	ThumbnailHeight int
}

func Load() *Config {
	var config Config

	// Get the current working directory
	basePath, err := os.Getwd()
	if err != nil {
		log.Fatal("❌ Error getting current working directory:", err)
	}

	// Detect environment (default: "development")
	config.AppEnv = getEnv("APP_ENV", "development")

	// Try to find the project root directory
	projectRoot := findProjectRoot(basePath)

	// Load .env file from project root
	envFile := filepath.Join(projectRoot, ".env")

	if err := godotenv.Load(envFile); err != nil {
		log.Printf("⚠ No .env file found at %s", envFile)
	} else {
		log.Printf("✅ Loaded .env from: %s", envFile)
	}

	// Server configuration
	config.Server.Port = getEnv("STORAGE_SERVICE_PORT", "8009")
	config.Server.ReadTimeout = time.Second * 30
	config.Server.WriteTimeout = time.Second * 30

	// Storage configuration
	config.Storage.UploadPath = getEnv("STORAGE_UPLOAD_PATH", "../../static-assets")
	config.Storage.BaseURL = getEnv("STORAGE_BASE_URL", "http://localhost:8009/static")
	config.Storage.MaxFileSize = 10 * 1024 * 1024 // 10MB
	config.Storage.AllowedTypes = []string{
		"image/jpeg", "image/jpg", "image/png", "image/gif", "image/webp",
		"video/mp4", "video/mov", "video/avi", "video/wmv", "video/flv",
		"application/pdf", "application/msword", "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	}
	config.Storage.ResizeWidth = 800
	config.Storage.ResizeHeight = 600
	config.Storage.ThumbnailWidth = 200
	config.Storage.ThumbnailHeight = 200

	return &config
}

// findProjectRoot tries to find the project root by looking for .env file
func findProjectRoot(startPath string) string {
	currentPath := startPath

	// Try up to 5 levels up to find the project root
	for i := 0; i < 5; i++ {
		// Check if .env exists in current path
		envPath := filepath.Join(currentPath, ".env")
		if _, err := os.Stat(envPath); err == nil {
			return currentPath
		}

		// Go up one level
		currentPath = filepath.Dir(currentPath)
	}

	// If not found, return the original path
	return startPath
}

// getEnv retrieves environment variable or returns a fallback value
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// Helper methods
func (c *Config) IsDevelopment() bool {
	return c.AppEnv == "development"
}

func (c *Config) IsProduction() bool {
	return c.AppEnv == "production"
}