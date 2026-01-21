package database

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/config"
	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLConnection(cfg config.DatabaseConfig) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	
	log.Println("Database connected successfully")
	return db
}

func RunMigrations(db *sql.DB) {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id BIGINT AUTO_INCREMENT PRIMARY KEY,
		email VARCHAR(255) NOT NULL UNIQUE,
		username VARCHAR(255) NOT NULL UNIQUE,
		password_hash VARCHAR(255) NOT NULL,
		role VARCHAR(50) NOT NULL DEFAULT 'user',
		is_active BOOLEAN NOT NULL DEFAULT TRUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		INDEX idx_email (email),
		INDEX idx_username (username)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
	`
	
	if _, err := db.Exec(createTableQuery); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	
	log.Println("Migrations completed successfully")
}