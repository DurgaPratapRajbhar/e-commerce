package models

import (
	"time"
)

type UserProfile struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `gorm:"unique;not null" json:"user_id" validate:"required"`
	FirstName   string    `gorm:"size:100" json:"first_name"`
	LastName    string    `gorm:"size:100" json:"last_name"`
	DateOfBirth *time.Time `gorm:"" json:"date_of_birth,omitempty"`
	Gender      string    `gorm:"size:20" json:"gender" validate:"oneof=male female other"`
	AvatarURL   string    `gorm:"size:500" json:"avatar_url"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (UserProfile) TableName() string {
	return "user_profiles"
}