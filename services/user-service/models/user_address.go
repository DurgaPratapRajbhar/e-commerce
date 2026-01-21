package models

import (
	"time"
)

type UserAddress struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	UserID       uint      `gorm:"not null" json:"user_id" validate:"required"`
	AddressType  string    `gorm:"size:20;not null" json:"address_type" validate:"oneof=home work other"`
	StreetAddress string   `gorm:"size:255;not null" json:"street_address"`
	City         string    `gorm:"size:100;not null" json:"city"`
	State        string    `gorm:"size:100;not null" json:"state"`
	PostalCode   string    `gorm:"size:20;not null" json:"postal_code"`
	Country      string    `gorm:"size:100;not null" json:"country"`
	IsDefault    bool      `gorm:"default:false" json:"is_default"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (UserAddress) TableName() string {
	return "user_addresses"
}