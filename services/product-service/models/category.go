package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Category struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	ParentID    uint64 `gorm:"index;not null;default:0" json:"parent_id"`
	Name        string `gorm:"type:varchar(255);not null" json:"name" validate:"required,min=3,max=255"`
	Slug        string `gorm:"type:varchar(255);unique;not null" json:"slug"`
	Description string `gorm:"type:text" json:"description" validate:"required,min=5"`
	Image       string `gorm:"type:varchar(255);null" json:"image"`
	Status      string `gorm:"type:enum('active', 'inactive');default:'active';not null" json:"status" validate:"oneof=active inactive"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Self-referencing relationship
	ParentCategory *Category  `gorm:"foreignKey:ParentID;references:ID;constraint:OnDelete:SET NULL" json:"parent_category,omitempty"`
	Children       []Category `gorm:"foreignKey:ParentID;references:ID" json:"children,omitempty"`
}

// TableName overrides the default table name
func (Category) TableName() string {
	return "categories"
}

// Validate performs custom validation for the Category model.
func (category *Category) Validate() error {
	validate := validator.New()
	return validate.Struct(category)
}

// BeforeCreate hook to set timestamps
func (category *Category) BeforeCreate(tx *gorm.DB) (err error) {
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()
	return
}

// BeforeUpdate hook to set timestamps
func (category *Category) BeforeUpdate(tx *gorm.DB) (err error) {
	category.UpdatedAt = time.Now()
	return
}
