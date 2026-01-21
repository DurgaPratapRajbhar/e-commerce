package models

import (
	"time"
)

type Cart struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint64    `gorm:"index;not null" json:"user_id"`
	ProductID uint64    `gorm:"index;not null" json:"product_id"`
	VariantID *uint64   `gorm:"index" json:"variant_id,omitempty"`
	Quantity  int       `gorm:"type:int;not null;default:1" json:"quantity"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}