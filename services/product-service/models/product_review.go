package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductReview struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID uint64    `gorm:"index;not null" json:"product_id"`                                // Foreign key to the product
	UserID    uint64    `gorm:"index;not null" json:"user_id"`                                   // Foreign key to the user
	Rating    int       `gorm:"type:int;not null" json:"rating" validate:"required,min=1,max=5"` // Rating between 1 and 5
	Review    string    `gorm:"type:text" json:"review" validate:"required,min=10"`              // Review text (min 10 chars)
	CreatedAt time.Time `json:"created_at"`                                                      // Timestamp for review creation

	// Relationships
	//Product Product `gorm:"foreignKey:ProductID;references:ID" json:"product"`
	// User    User    `gorm:"foreignKey:UserID;references:ID" json:"user"`
}

// TableName overrides the default table name for the ProductReview model
func (ProductReview) TableName() string {
	return "product_reviews"
}

// BeforeCreate hook to set timestamps
func (productReview *ProductReview) BeforeCreate(tx *gorm.DB) (err error) {
	productReview.CreatedAt = time.Now()
	return
}
