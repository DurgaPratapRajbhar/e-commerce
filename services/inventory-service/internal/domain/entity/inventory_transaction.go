package entity

import (
	"time"
)

type InventoryTransaction struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	ProductID      uint      `json:"product_id" gorm:"not null;index" validate:"required"`
	VariantID      uint      `json:"variant_id" gorm:"not null;index" validate:"required"`
	TransactionType string   `json:"transaction_type" gorm:"not null;size:20" validate:"required,oneof=in out reserved" example:"in"`
	Quantity       int       `json:"quantity" gorm:"not null" validate:"min=1" example:"10"`
	ReferenceID    *uint     `json:"reference_id,omitempty" gorm:"index" example:"123"`
	Notes          string    `json:"notes,omitempty" gorm:"size:500" validate:"max=500"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
}

// InventoryTransactionRequest represents the request to create an inventory transaction
type InventoryTransactionRequest struct {
	ProductID       uint   `json:"product_id" validate:"required"`
	VariantID       uint   `json:"variant_id" validate:"required"`
	TransactionType string `json:"transaction_type" validate:"required,oneof=in out reserved"`
	Quantity        int    `json:"quantity" validate:"min=1"`
	ReferenceID     *uint  `json:"reference_id,omitempty"`
	Notes           string `json:"notes,omitempty" validate:"max=500"`
}