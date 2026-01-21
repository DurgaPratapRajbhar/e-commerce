package entity

import (
	"time"
)

type TrackingEvent struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	ShipmentID  uint      `json:"shipment_id" gorm:"not null;index" validate:"required"`
	EventType   string    `json:"event_type" gorm:"size:50;not null" validate:"required,max=50"`
	Location    string    `json:"location" gorm:"size:200" validate:"max=200"`
	Description string    `json:"description" gorm:"size:500" validate:"max=500"`
	Timestamp   time.Time `json:"timestamp" gorm:"not null;autoCreateTime"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
}

// TrackingEventCreateRequest represents the request to create a tracking event
type TrackingEventCreateRequest struct {
	ShipmentID  uint      `json:"shipment_id" validate:"required"`
	EventType   string    `json:"event_type" validate:"required,max=50"`
	Location    string    `json:"location,omitempty" validate:"max=200"`
	Description string    `json:"description,omitempty" validate:"max=500"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
}