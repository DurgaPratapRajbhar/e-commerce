package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

// Main Config Structure
type Config struct {
	AppEnv       string
	ImageGallery string
	Server       ServerConfig
	Database     DatabaseConfig
	Services     ServicesConfig
	Auth         AuthConfig
	JWT          JWTConfig
	RateLimit    RateLimitConfig
	External     ExternalConfig
	Cors         CorsConfig
	Storage      StorageConfig
}

// Server Configuration
type ServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// Database Configuration
type DatabaseConfig struct {
	Type        string
	Host        string
	Port        string
	User        string
	Password    string
	DBName      string
	UserDB      string
	ProductDB   string
	OrderDB     string
	CartDB      string
	PaymentDB   string
	InventoryDB string
	ShippingDB  string
	Charset     string
	ParseTime   bool
	Loc         string
	AuthDSN     string
}

// Services Configuration
type ServicesConfig struct {
	AuthService      ServiceConfig
	UserService      ServiceConfig
	ProductService   ServiceConfig
	CartService      ServiceConfig
	OrderService     ServiceConfig
	PaymentService   ServiceConfig
	ShippingService  ServiceConfig
	InventoryService ServiceConfig
	StorageService   ServiceConfig
	LoggingService   ServiceConfig
}

type ServiceConfig struct {
	URL     string
	Port    string
	Timeout time.Duration
}

// Auth Configuration (Gateway JWT handling)
type AuthConfig struct {
	JWTSecret   string
	TokenExpiry time.Duration
}

// JWT Configuration (Service-level)
type JWTConfig struct {
	SecretKey   string
	ExpiryHours int
}

// Rate Limiting Configuration
type RateLimitConfig struct {
	RequestsPerMinute int
	BurstSize         int
}

// External Services Configuration
type ExternalConfig struct {
	SMTPHost     string
	SMTPPort     string
	SMTPUser     string
	SMTPPassword string
	SMTPFrom     string
}

// CORS Configuration
type CorsConfig struct {
	CorsOrigins []string
}

// Storage Configuration
type StorageConfig struct {
	ImagePath string
}

// Load loads the unified configuration
func Load() (*Config, error) {
	// Find project root and load .env
	basePath, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("error getting working directory: %w", err)
	}

	projectRoot := findProjectRoot(basePath)
	envFile := filepath.Join(projectRoot, ".env")

	if err := godotenv.Load(envFile); err != nil {
		log.Printf("⚠️  No .env file found at %s, using defaults", envFile)
	} else {
		log.Printf("✅ Loaded .env from: %s", envFile)
	}

	// Detect environment
	appEnv := getEnv("APP_ENV", "development")

	cfg := &Config{
		AppEnv:       appEnv,
		ImageGallery: getEnv("IMAGE_GALLERY", "../../image_gallery/"),

		// Server Configuration
		Server: ServerConfig{
			Port:         getEnv("GATEWAY_PORT", getEnv("SERVER_PORT", "8080")),
			ReadTimeout:  getDuration("GATEWAY_READ_TIMEOUT", 30),
			WriteTimeout: getDuration("GATEWAY_WRITE_TIMEOUT", 30),
		},

		// Database Configuration
		Database: DatabaseConfig{
			Type:        getEnv("DB_TYPE", "mysql"),
			Host:        getEnv("DB_HOST", "localhost"),
			Port:        getEnv("DB_PORT", "3306"),
			User:        getEnv("DB_USER", "root"),
			Password:    getEnv("DB_PASSWORD", "password"),
			DBName:      getEnv("AUTH_DB", "auth_service_db"),
			UserDB:      getEnv("USER_DB", "user_service_db"),
			ProductDB:   getEnv("PRODUCT_DB", "product_service_db"),
			OrderDB:     getEnv("ORDER_DB", "order_service_db"),
			CartDB:      getEnv("CART_DB", "cart_service_db"),
			PaymentDB:   getEnv("PAYMENT_DB", "payment_service_db"),
			InventoryDB: getEnv("INVENTORY_DB", "inventory_service_db"),
			ShippingDB:  getEnv("SHIPPING_DB", "shipping_service_db"),
			Charset:     getEnv("DB_CHARSET", "utf8mb4"),
			ParseTime:   getEnvBool("DB_PARSE_TIME", true),
			Loc:         getEnv("DB_LOC", "Local"),
			AuthDSN:     getEnv("AUTH_DB_DSN", ""),
		},

		// Services Configuration (combining both approaches)
		Services: ServicesConfig{
			AuthService: ServiceConfig{
				URL:     getEnv("AUTH_SERVICE_URL", "http://localhost:8081"),
				Port:    getEnv("AUTH_PORT", "8081"),
				Timeout: getDuration("AUTH_SERVICE_TIMEOUT", 30),
			},
			UserService: ServiceConfig{
				URL:     getEnv("USER_SERVICE_URL", "http://localhost:8082"),
				Port:    getEnv("USER_PORT", "8082"),
				Timeout: getDuration("USER_SERVICE_TIMEOUT", 30),
			},
			ProductService: ServiceConfig{
				URL:     getEnv("PRODUCT_SERVICE_URL", "http://localhost:8083"),
				Port:    getEnv("PRODUCT_PORT", "8083"),
				Timeout: getDuration("PRODUCT_SERVICE_TIMEOUT", 30),
			},
			CartService: ServiceConfig{
				URL:     getEnv("CART_SERVICE_URL", "http://localhost:8084"),
				Port:    getEnv("CART_PORT", "8084"),
				Timeout: getDuration("CART_SERVICE_TIMEOUT", 30),
			},
			OrderService: ServiceConfig{
				URL:     getEnv("ORDER_SERVICE_URL", "http://localhost:8085"),
				Port:    getEnv("ORDER_PORT", "8085"),
				Timeout: getDuration("ORDER_SERVICE_TIMEOUT", 30),
			},
			PaymentService: ServiceConfig{
				URL:     getEnv("PAYMENT_SERVICE_URL", "http://localhost:8086"),
				Port:    getEnv("PAYMENT_PORT", "8086"),
				Timeout: getDuration("PAYMENT_SERVICE_TIMEOUT", 30),
			},
			ShippingService: ServiceConfig{
				URL:     getEnv("SHIPPING_SERVICE_URL", "http://localhost:8088"),
				Port:    getEnv("SHIPPING_PORT", "8088"),
				Timeout: getDuration("SHIPPING_SERVICE_TIMEOUT", 30),
			},
			InventoryService: ServiceConfig{
				URL:     getEnv("INVENTORY_SERVICE_URL", "http://localhost:8087"),
				Port:    getEnv("INVENTORY_PORT", "8087"),
				Timeout: getDuration("INVENTORY_SERVICE_TIMEOUT", 30),
			},
			StorageService: ServiceConfig{
				URL:     getEnv("STORAGE_SERVICE_URL", "http://localhost:8089"),
				Port:    getEnv("STORAGE_PORT", "8089"),
				Timeout: getDuration("STORAGE_SERVICE_TIMEOUT", 30),
			},
			LoggingService: ServiceConfig{
				URL:     getEnv("LOGGING_SERVICE_URL", "http://localhost:8090"),
				Port:    getEnv("LOGGING_PORT", "8090"),
				Timeout: getDuration("LOGGING_SERVICE_TIMEOUT", 30),
			},
		},

		// Auth Configuration (for Gateway)
		Auth: AuthConfig{
			JWTSecret:   getEnv("JWT_SECRET", "your-secret-key-here"),
			TokenExpiry: getDuration("JWT_TOKEN_EXPIRY", 86400), // 24 hours
		},

		// JWT Configuration (for individual services)
		JWT: JWTConfig{
			SecretKey:   getEnv("JWT_SECRET", "change-me"),
			ExpiryHours: getEnvInt("JWT_EXPIRY_HOURS", 24),
		},

		// Rate Limiting
		RateLimit: RateLimitConfig{
			RequestsPerMinute: getEnvInt("RATE_LIMIT_RPM", 100),
			BurstSize:         getEnvInt("RATE_LIMIT_BURST", 20),
		},

		// External Services
		External: ExternalConfig{
			SMTPHost:     getEnv("SMTP_HOST", "smtp.gmail.com"),
			SMTPPort:     getEnv("SMTP_PORT", "587"),
			SMTPUser:     getEnv("SMTP_USER", ""),
			SMTPPassword: getEnv("SMTP_PASSWORD", ""),
			SMTPFrom:     getEnv("SMTP_FROM", "noreply@example.com"),
		},

		// CORS Configuration
		Cors: CorsConfig{
			CorsOrigins: []string{
				getEnv("CORS_ORIGIN_1", "http://localhost:3000"),
				getEnv("CORS_ORIGIN_2", "http://localhost:5173"),
			},
		},

		// Storage Configuration
		Storage: StorageConfig{
			ImagePath: getEnv("IMAGE_PATH", getEnv("IMAGE_GALLERY", "../../image_gallery/")),
		},
	}

	// Debug: Print loaded config
	printConfig(cfg)

	return cfg, nil
}

// LoadConfig is an alias for Load() for backward compatibility
func LoadConfig() Config {
	cfg, err := Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}
	return *cfg
}

// Helper Functions
func findProjectRoot(startPath string) string {
	currentPath := startPath
	for i := 0; i < 5; i++ {
		envPath := filepath.Join(currentPath, ".env")
		if _, err := os.Stat(envPath); err == nil {
			return currentPath
		}
		currentPath = filepath.Dir(currentPath)
	}
	return startPath
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if value, exists := os.LookupEnv(key); exists {
		var result int
		if _, err := fmt.Sscanf(value, "%d", &result); err == nil {
			return result
		}
	}
	return fallback
}

func getEnvBool(key string, fallback bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		return strings.ToLower(value) == "true" || value == "1"
	}
	return fallback
}

func getDuration(key string, fallbackSeconds int) time.Duration {
	if value, exists := os.LookupEnv(key); exists {
		var seconds int
		if _, err := fmt.Sscanf(value, "%d", &seconds); err == nil {
			return time.Duration(seconds) * time.Second
		}
	}
	return time.Duration(fallbackSeconds) * time.Second
}

func maskSecret(secret string) string {
	if len(secret) <= 8 {
		return "****"
	}
	return secret[:4] + "****" + secret[len(secret)-4:]
}

func printConfig(cfg *Config) {
	log.Printf("✅ Configuration Loaded:")
	log.Printf("   Environment: %s", cfg.AppEnv)
	log.Printf("   Server Port: %s", cfg.Server.Port)
	log.Printf("   JWT Secret: %s", maskSecret(cfg.Auth.JWTSecret))
	log.Printf("   Database Host: %s:%s", cfg.Database.Host, cfg.Database.Port)
	log.Printf("   Auth Service: %s (Port: %s)", cfg.Services.AuthService.URL, cfg.Services.AuthService.Port)
	log.Printf("   User Service: %s (Port: %s)", cfg.Services.UserService.URL, cfg.Services.UserService.Port)
	log.Printf("   Product Service: %s (Port: %s)", cfg.Services.ProductService.URL, cfg.Services.ProductService.Port)
	log.Printf("   Cart Service: %s (Port: %s)", cfg.Services.CartService.URL, cfg.Services.CartService.Port)
	log.Printf("   Order Service: %s (Port: %s)", cfg.Services.OrderService.URL, cfg.Services.OrderService.Port)
	log.Printf("   Payment Service: %s (Port: %s)", cfg.Services.PaymentService.URL, cfg.Services.PaymentService.Port)
	log.Printf("   Shipping Service: %s (Port: %s)", cfg.Services.ShippingService.URL, cfg.Services.ShippingService.Port)
	log.Printf("   Inventory Service: %s (Port: %s)", cfg.Services.InventoryService.URL, cfg.Services.InventoryService.Port)
}