package domain

import (
	"context"
	"github.com/DurgaPratapRajbhar/e-commerce/shipping-service/internal/domain/entity"
)

type ShipmentRepository interface {
	Create(ctx context.Context, shipment *entity.Shipment) error
	GetByID(ctx context.Context, id uint) (*entity.Shipment, error)
	GetByOrderID(ctx context.Context, orderID uint) (*entity.Shipment, error)
	GetByTrackingNumber(ctx context.Context, trackingNumber string) (*entity.Shipment, error)
	Update(ctx context.Context, shipment *entity.Shipment) error
	UpdateStatus(ctx context.Context, id uint, status string) error
	Delete(ctx context.Context, id uint) error
	GetByStatus(ctx context.Context, status string) ([]entity.Shipment, error)
	GetAll(ctx context.Context, limit, offset int) ([]entity.Shipment, error)
}

type TrackingEventRepository interface {
	Create(ctx context.Context, event *entity.TrackingEvent) error
	GetByShipmentID(ctx context.Context, shipmentID uint) ([]entity.TrackingEvent, error)
	GetLatestByShipmentID(ctx context.Context, shipmentID uint) (*entity.TrackingEvent, error)
	GetByEventType(ctx context.Context, eventType string) ([]entity.TrackingEvent, error)
	GetAll(ctx context.Context, limit, offset int) ([]entity.TrackingEvent, error)
}