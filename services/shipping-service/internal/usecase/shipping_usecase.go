package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"
	"github.com/DurgaPratapRajbhar/e-commerce/shipping-service/internal/domain"
	"github.com/DurgaPratapRajbhar/e-commerce/shipping-service/internal/domain/entity"
)

type ShippingUseCase struct {
	shipmentRepo   domain.ShipmentRepository
	trackingRepo   domain.TrackingEventRepository
}

func NewShippingUseCase(
	shipmentRepo domain.ShipmentRepository,
	trackingRepo domain.TrackingEventRepository,
) *ShippingUseCase {
	return &ShippingUseCase{
		shipmentRepo:   shipmentRepo,
		trackingRepo:   trackingRepo,
	}
}

func (uc *ShippingUseCase) CreateShipment(ctx context.Context, req *entity.ShipmentCreateRequest) (*entity.Shipment, error) {
	// Check if a shipment already exists for this order
	existing, err := uc.shipmentRepo.GetByOrderID(ctx, req.OrderID)
	if err == nil && existing != nil {
		return nil, errors.New("shipment already exists for this order")
	}

	// Generate tracking number if not provided
	trackingNumber := req.Carrier + fmt.Sprintf("%d", time.Now().Unix())
	if req.Carrier == "" {
		trackingNumber = fmt.Sprintf("SHIP-%d", time.Now().Unix())
	}

	shipment := &entity.Shipment{
		OrderID:           req.OrderID,
		TrackingNumber:    trackingNumber,
		Carrier:           req.Carrier,
		ShippingMethod:    req.ShippingMethod,
		Status:            "pending",
		EstimatedDelivery: req.EstimatedDelivery,
	}

	err = uc.shipmentRepo.Create(ctx, shipment)
	if err != nil {
		return nil, fmt.Errorf("failed to create shipment: %w", err)
	}

	// Create initial tracking event
	initialEvent := &entity.TrackingEvent{
		ShipmentID:  shipment.ID,
		EventType:   "created",
		Location:    "Origin Facility",
		Description: "Shipment created and awaiting processing",
		Timestamp:   time.Now(),
	}
	
	err = uc.trackingRepo.Create(ctx, initialEvent)
	if err != nil {
		return nil, fmt.Errorf("failed to create initial tracking event: %w", err)
	}

	return shipment, nil
}

func (uc *ShippingUseCase) GetShipment(ctx context.Context, id uint) (*entity.Shipment, error) {
	return uc.shipmentRepo.GetByID(ctx, id)
}

func (uc *ShippingUseCase) GetShipmentByOrderID(ctx context.Context, orderID uint) (*entity.Shipment, error) {
	return uc.shipmentRepo.GetByOrderID(ctx, orderID)
}

func (uc *ShippingUseCase) GetShipmentByTrackingNumber(ctx context.Context, trackingNumber string) (*entity.Shipment, error) {
	return uc.shipmentRepo.GetByTrackingNumber(ctx, trackingNumber)
}

func (uc *ShippingUseCase) UpdateShipment(ctx context.Context, id uint, req *entity.ShipmentUpdateRequest) (*entity.Shipment, error) {
	shipment, err := uc.shipmentRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided in request
	if req.TrackingNumber != nil {
		shipment.TrackingNumber = *req.TrackingNumber
	}
	if req.Carrier != nil {
		shipment.Carrier = *req.Carrier
	}
	if req.ShippingMethod != nil {
		shipment.ShippingMethod = *req.ShippingMethod
	}
	if req.Status != nil {
		shipment.Status = *req.Status
	}
	if req.EstimatedDelivery != nil {
		shipment.EstimatedDelivery = req.EstimatedDelivery
	}
	if req.ActualDelivery != nil {
		shipment.ActualDelivery = req.ActualDelivery
	}

	err = uc.shipmentRepo.Update(ctx, shipment)
	if err != nil {
		return nil, fmt.Errorf("failed to update shipment: %w", err)
	}

	return shipment, nil
}

func (uc *ShippingUseCase) UpdateShipmentStatus(ctx context.Context, id uint, status string) error {
	// Validate status
	validStatuses := map[string]bool{
		"pending":      true,
		"in_transit":   true,
		"delivered":    true,
		"returned":     true,
	}
	if !validStatuses[status] {
		return errors.New("invalid shipment status")
	}

	err := uc.shipmentRepo.UpdateStatus(ctx, id, status)
	if err != nil {
		return fmt.Errorf("failed to update shipment status: %w", err)
	}

	return nil
}

func (uc *ShippingUseCase) DeleteShipment(ctx context.Context, id uint) error {
	return uc.shipmentRepo.Delete(ctx, id)
}

func (uc *ShippingUseCase) GetShipmentsByStatus(ctx context.Context, status string) ([]entity.Shipment, error) {
	return uc.shipmentRepo.GetByStatus(ctx, status)
}

func (uc *ShippingUseCase) GetShipments(ctx context.Context, limit, offset int) ([]entity.Shipment, error) {
	return uc.shipmentRepo.GetAll(ctx, limit, offset)
}

func (uc *ShippingUseCase) CreateTrackingEvent(ctx context.Context, req *entity.TrackingEventCreateRequest) error {
	event := &entity.TrackingEvent{
		ShipmentID:  req.ShipmentID,
		EventType:   req.EventType,
		Location:    req.Location,
		Description: req.Description,
		Timestamp:   req.Timestamp,
	}

	if event.Timestamp.IsZero() {
		event.Timestamp = time.Now()
	}

	err := uc.trackingRepo.Create(ctx, event)
	if err != nil {
		return fmt.Errorf("failed to create tracking event: %w", err)
	}

	return nil
}

func (uc *ShippingUseCase) GetTrackingEvents(ctx context.Context, shipmentID uint) ([]entity.TrackingEvent, error) {
	return uc.trackingRepo.GetByShipmentID(ctx, shipmentID)
}

func (uc *ShippingUseCase) GetLatestTrackingEvent(ctx context.Context, shipmentID uint) (*entity.TrackingEvent, error) {
	return uc.trackingRepo.GetLatestByShipmentID(ctx, shipmentID)
}