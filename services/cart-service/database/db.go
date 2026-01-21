package database

import (
	"fmt"
	"log"

	"github.com/DurgaPratapRajbhar/e-commerce/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(config config.Config) {
	var err error
	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.CartDB,
		config.Database.Charset,
		config.Database.ParseTime,
		config.Database.Loc)
	
	fmt.Printf("CartDB connection string: %s\n", dsn)
	 
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	fmt.Println("Connected to database successfully")
}