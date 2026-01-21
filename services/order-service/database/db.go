package database

import (
	"context"
	"fmt"
	"time"

	"github.com/DurgaPratapRajbhar/e-commerce/pkg/logger"
	"github.com/DurgaPratapRajbhar/e-commerce/services/order-service/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(connString string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
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

	if err := sqlDB.Ping(); err != nil {
		logger.Logger.Error("Database connection is not alive:", err)
		return nil, fmt.Errorf("database connection is not alive: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	logger.Logger.Info("Connected to MySQL successfully.")
	return db, nil
}

func Migrate(db *gorm.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := db.WithContext(ctx).AutoMigrate(&models.Order{}, &models.OrderItem{}); err != nil {
		logger.Logger.Error("Error migrating MySQL database:", err)
		return fmt.Errorf("error migrating MySQL database: %w", err)
	}

	logger.Logger.Info("Database migrated successfully.")
	return nil
}
