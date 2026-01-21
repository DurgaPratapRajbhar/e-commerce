package entity

import (
	"time"
)

type Inventory struct {
	ID               uint      `json:"id" gorm:"primaryKey"`
	ProductID        uint      `json:"product_id" gorm:"not null;index" validate:"required"`
	VariantID        uint      `json:"variant_id" gorm:"not null;index" validate:"required"`
	Quantity         int       `json:"quantity" gorm:"not null;default:0" validate:"min=0"`
	ReservedQuantity int       `json:"reserved_quantity" gorm:"not null;default:0" validate:"min=0"`
	WarehouseLocation string   `json:"warehouse_location" gorm:"size:255" validate:"max=255"`
	LastUpdated      time.Time `json:"last_updated" gorm:"autoUpdateTime"`
	CreatedAt        time.Time `json:"created_at" gorm:"autoCreateTime"`
}

// InventoryUpdateRequest represents the request to update inventory
type InventoryUpdateRequest struct {
	ProductID         uint `json:"product_id" validate:"required"`
	VariantID         uint `json:"variant_id" validate:"required"`
	QuantityChange    int  `json:"quantity_change" validate:"required,ne=0"`
	TransactionType   string `json:"transaction_type" validate:"required,oneof=in out reserved"`
	ReferenceID       *uint `json:"reference_id,omitempty"`
	WarehouseLocation *string `json:"warehouse_location,omitempty"`
}