package mysql

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"github.com/DurgaPratapRajbhar/e-commerce/shipping-service/internal/domain/entity"
	"github.com/DurgaPratapRajbhar/e-commerce/shipping-service/internal/domain"
)

type trackingEventRepository struct {
	db *gorm.DB
}

func NewTrackingEventRepository(db *gorm.DB) domain.TrackingEventRepository {
	return &trackingEventRepository{db: db}
}

func (r *trackingEventRepository) Create(ctx context.Context, event *entity.TrackingEvent) error {
	return r.db.WithContext(ctx).Create(event).Error
}

func (r *trackingEventRepository) GetByShipmentID(ctx context.Context, shipmentID uint) ([]entity.TrackingEvent, error) {
	var events []entity.TrackingEvent
	err := r.db.WithContext(ctx).Where("shipment_id = ?", shipmentID).Order("timestamp ASC").Find(&events).Error
	return events, err
}

func (r *trackingEventRepository) GetLatestByShipmentID(ctx context.Context, shipmentID uint) (*entity.TrackingEvent, error) {
	var event entity.TrackingEvent
	err := r.db.WithContext(ctx).Where("shipment_id = ?", shipmentID).Order("timestamp DESC").First(&event).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("tracking event not found")
		}
		return nil, err
	}
	return &event, nil
}

func (r *trackingEventRepository) GetByEventType(ctx context.Context, eventType string) ([]entity.TrackingEvent, error) {
	var events []entity.TrackingEvent
	err := r.db.WithContext(ctx).Where("event_type = ?", eventType).Find(&events).Error
	return events, err
}

func (r *trackingEventRepository) GetAll(ctx context.Context, limit, offset int) ([]entity.TrackingEvent, error) {
	var events []entity.TrackingEvent
	err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&events).Error
	return events, err
}