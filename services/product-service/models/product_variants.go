package models

import (
	"time"
)

type ProductVariantTest struct {
	ID              uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID       uint64    `gorm:"index;not null" json:"product_id"` // Foreign key to the product
	AttributeName   string    `gorm:"type:varchar(100);not null" json:"attribute_name" validate:"required,max=100"`
	AttributeValue  string    `gorm:"type:varchar(100);not null" json:"attribute_value" validate:"required,max=100"`
	AdditionalPrice float64   `gorm:"type:decimal(10,2);default:0" json:"additional_price"` // Extra price for this variant
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`

	//Product Product `gorm:"foreignKey:ProductID;references:ID" json:"product"`
}

// func (ProductVariantTest) TableName() string {
// 	return "product_variants"
// }

// func (productVariant *ProductVariantTest) BeforeCreate(tx *gorm.DB) (err error) {
// 	ProductVariantTest.CreatedAt = time.Now()
// 	ProductVariantTest.UpdatedAt = time.Now()
// 	return
// }

// func (productVariant *ProductVariant) BeforeUpdate(tx *gorm.DB) (err error) {
// 	productVariant.UpdatedAt = time.Now()
// 	return
// }

// id	product_id	attribute_name	attribute_value	additional_price
// 1	1	Color	Red	0
// 2	1	Color	Blue	500
// 3	1	Storage	128GB	0
// 4	1	Storage	256GB	5000
