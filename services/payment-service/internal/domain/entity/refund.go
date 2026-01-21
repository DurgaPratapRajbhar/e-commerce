package entity

import (
	"time"
)

type Refund struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	PaymentID   uint      `json:"payment_id" gorm:"not null;index" validate:"required"`
	OrderID     uint      `json:"order_id" gorm:"not null;index" validate:"required"`
	Amount      float64   `json:"amount" gorm:"not null" validate:"required,gt=0"`
	Reason      string    `json:"reason" gorm:"size:200" validate:"max=200"`
	Status      string    `json:"status" gorm:"size:20;default:'pending'" validate:"oneof=pending processed failed" example:"pending"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
}

// RefundCreateRequest represents the request to create a refund
type RefundCreateRequest struct {
	PaymentID uint      `json:"payment_id" validate:"required"`
	OrderID   uint      `json:"order_id" validate:"required"`
	Amount    float64   `json:"amount" validate:"required,gt=0"`
	Reason    string    `json:"reason" validate:"max=200"`
}