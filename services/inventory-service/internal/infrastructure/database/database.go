package database

import (
	"fmt"
	"log"

	"github.com/DurgaPratapRajbhar/e-commerce/pkg/config"
	"github.com/joho/godotenv"
	"github.com/DurgaPratapRajbhar/e-commerce/inventory-service/internal/domain/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewMySQLConnection() *gorm.DB {
	// Load environment variables from .env file
	godotenv.Load()

	// Load the shared config
	globalConfig := config.LoadConfig()

	// Construct DSN from shared config
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=%s&parseTime=%t&loc=%s",
		globalConfig.Database.User,
		globalConfig.Database.Password,
		globalConfig.Database.Host,
		globalConfig.Database.Port,
		globalConfig.Database.InventoryDB, // Use the inventory-specific DB name
		globalConfig.Database.Charset,
		globalConfig.Database.ParseTime,
		globalConfig.Database.Loc,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	return db
}

func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(
		&entity.Inventory{},
		&entity.InventoryTransaction{},
	)
	if err != nil {
		log.Fatal("Failed to run migrations:", err)
	}
	fmt.Println("Database migrations completed successfully")
}