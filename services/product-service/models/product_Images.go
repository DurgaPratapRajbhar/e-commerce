package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductImage struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID uint64    `gorm:"index;not null" json:"product_id"`
	ImageURL  string    `gorm:"type:varchar(255);not null" json:"image_url" validate:"required"` // Removed `url`
	AltText   string    `gorm:"type:varchar(255);not null" json:"alt_text" validate:"required,max=255"`
	IsPrimary bool      `gorm:"type:bool;default:false" json:"is_primary"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Product *Product `gorm:"foreignKey:ProductID;references:ID" json:"product,omitempty"` // Changed to pointer
}

func (ProductImage) TableName() string {
	return "product_images"
}

func (productImage *ProductImage) BeforeCreate(tx *gorm.DB) (err error) {
	productImage.CreatedAt = time.Now()
	productImage.UpdatedAt = time.Now()
	return
}

// BeforeUpdate hook to set timestamps
func (productImage *ProductImage) BeforeUpdate(tx *gorm.DB) (err error) {
	productImage.UpdatedAt = time.Now()
	return
}
