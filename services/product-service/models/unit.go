package models

import "time"

// UnitOfMeasurement model
type UnitOfMeasurement struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"type:varchar(50)"`
	Symbol      string    `gorm:"type:varchar(10)"  json:"symbol"`
	Description string    `gorm:"type:varchar(255)" json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (UnitOfMeasurement) TableName() string { return "units_of_measurement" }
