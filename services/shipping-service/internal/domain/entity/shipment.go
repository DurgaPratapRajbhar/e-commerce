package entity

import (
	"time"
)

type Shipment struct {
	ID                 uint      `json:"id" gorm:"primaryKey"`
	OrderID            uint      `json:"order_id" gorm:"not null;index" validate:"required"`
	TrackingNumber     string    `json:"tracking_number" gorm:"unique;not null;size:100" validate:"required,max=100"`
	Carrier            string    `json:"carrier" gorm:"size:100" validate:"max=100"`
	ShippingMethod     string    `json:"shipping_method" gorm:"size:50" validate:"max=50"`
	Status             string    `json:"status" gorm:"size:20;default:'pending'" validate:"oneof=pending in_transit delivered returned" example:"pending"`
	EstimatedDelivery  *time.Time `json:"estimated_delivery,omitempty"`
	ActualDelivery     *time.Time `json:"actual_delivery,omitempty"`
	CreatedAt          time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt          time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// ShipmentCreateRequest represents the request to create a shipment
type ShipmentCreateRequest struct {
	OrderID           uint      `json:"order_id" validate:"required"`
	Carrier           string    `json:"carrier" validate:"max=100"`
	ShippingMethod    string    `json:"shipping_method" validate:"max=50"`
	EstimatedDelivery *time.Time `json:"estimated_delivery,omitempty"`
}

// ShipmentUpdateRequest represents the request to update a shipment
type ShipmentUpdateRequest struct {
	TrackingNumber    *string    `json:"tracking_number,omitempty" validate:"omitempty,max=100"`
	Carrier           *string    `json:"carrier,omitempty" validate:"omitempty,max=100"`
	ShippingMethod    *string    `json:"shipping_method,omitempty" validate:"omitempty,max=50"`
	Status            *string    `json:"status,omitempty" validate:"omitempty,oneof=pending in_transit delivered returned"`
	EstimatedDelivery *time.Time `json:"estimated_delivery,omitempty"`
	ActualDelivery    *time.Time `json:"actual_delivery,omitempty"`
}