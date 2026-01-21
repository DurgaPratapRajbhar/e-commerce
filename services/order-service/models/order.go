package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	UserID        uint           `json:"user_id"`
	TotalAmount   float64        `json:"total_amount"`
	Status        string         `json:"status"`
	PaymentStatus string         `json:"payment_status"`
	PaymentID     *string        `json:"payment_id"`
	ShippingID    uint           `json:"shipping_id"`
	OrderItems    []OrderItem    `json:"order_items" gorm:"foreignKey:OrderID"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type OrderItem struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	OrderID   uint           `json:"order_id"`
	ProductID uint           `json:"product_id"`
	Quantity  int            `json:"quantity"`
	Price     float64        `json:"price"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
