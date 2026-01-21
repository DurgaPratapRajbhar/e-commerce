package entity

import (
	"time"
)

type Payment struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	OrderID         uint      `json:"order_id" gorm:"not null;index" validate:"required"`
	UserID          uint      `json:"user_id" gorm:"not null;index" validate:"required"`
	Amount          float64   `json:"amount" gorm:"not null" validate:"required,gt=0"`
	PaymentMethod   string    `json:"payment_method" gorm:"size:20;not null" validate:"required,oneof=card upi wallet cod" example:"card"`
	PaymentStatus   string    `json:"payment_status" gorm:"size:20;default:'pending'" validate:"oneof=pending success failed refunded" example:"pending"`
	TransactionID   string    `json:"transaction_id" gorm:"size:100" validate:"max=100"`
	GatewayResponse *string    `json:"gateway_response,omitempty" gorm:"type:text" validate:"max=1000"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// PaymentCreateRequest represents the request to create a payment
type PaymentCreateRequest struct {
	OrderID       uint      `json:"order_id" validate:"required"`
	UserID        uint      `json:"user_id" validate:"required"`
	Amount        float64   `json:"amount" validate:"required,gt=0"`
	PaymentMethod string    `json:"payment_method" validate:"required,oneof=card upi wallet cod"`
}

// PaymentUpdateRequest represents the request to update a payment
type PaymentUpdateRequest struct {
	PaymentStatus   *string `json:"payment_status,omitempty" validate:"omitempty,oneof=pending success failed refunded"`
	TransactionID   *string `json:"transaction_id,omitempty" validate:"omitempty,max=100"`
	GatewayResponse *string `json:"gateway_response,omitempty" validate:"omitempty,max=1000"`
}