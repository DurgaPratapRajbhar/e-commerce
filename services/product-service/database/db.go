package database

import (
	"context"
	"fmt"
	"time"

	"github.com/DurgaPratapRajbhar/e-commerce/pkg/logger"

	"github.com/DurgaPratapRajbhar/e-commerce/product-service/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitDB initializes and returns a database connection
func InitDB(connString string) (*gorm.DB, error) {
	//fmt.Println(connString)
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // Prevent foreign key issues
	})
	if err != nil {
		logger.Logger.Error("Failed to connect to MySQL:", err)
		return nil, fmt.Errorf("failed to connect to MySQL: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Logger.Error("Failed to get DB instance:", err)
		return nil, fmt.Errorf("failed to get DB instance: %w", err)
	}

	// ✅ SQL Connection Health Check
	if err := sqlDB.Ping(); err != nil {
		logger.Logger.Error("Database connection is not alive:", err)
		return nil, fmt.Errorf("database connection is not alive: %w", err)
	}

	// ✅ Optimized Connection Pooling
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(30 * time.Minute) // Prevent stale connections

	logger.Logger.Info("Connected to MySQL successfully.")
	return db, nil
}

// Migrate applies database migrations with a timeout
func Migrate(db *gorm.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second) // Set 30s timeout
	defer cancel()

	if err := db.WithContext(ctx).AutoMigrate(
		&models.Category{},
		&models.Product{},
		&models.ProductImage{},
		&models.ProductReview{},
		&models.ProductVariant{},
		&models.UnitOfMeasurement{},
		&models.ProductAttribute{},
	); err != nil {
		logger.Logger.Error("Error migrating MySQL database:", err)
		return fmt.Errorf("error migrating MySQL database: %w", err)
	}

	logger.Logger.Info("Database migrated successfully.")
	return nil
}
