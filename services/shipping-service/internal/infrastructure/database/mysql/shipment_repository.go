package mysql

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"github.com/DurgaPratapRajbhar/e-commerce/shipping-service/internal/domain/entity"
	"github.com/DurgaPratapRajbhar/e-commerce/shipping-service/internal/domain"
)

type shipmentRepository struct {
	db *gorm.DB
}

func NewShipmentRepository(db *gorm.DB) domain.ShipmentRepository {
	return &shipmentRepository{db: db}
}

func (r *shipmentRepository) Create(ctx context.Context, shipment *entity.Shipment) error {
	return r.db.WithContext(ctx).Create(shipment).Error
}

func (r *shipmentRepository) GetByID(ctx context.Context, id uint) (*entity.Shipment, error) {
	var shipment entity.Shipment
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&shipment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("shipment not found")
		}
		return nil, err
	}
	return &shipment, nil
}

func (r *shipmentRepository) GetByOrderID(ctx context.Context, orderID uint) (*entity.Shipment, error) {
	var shipment entity.Shipment
	err := r.db.WithContext(ctx).Where("order_id = ?", orderID).First(&shipment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("shipment not found")
		}
		return nil, err
	}
	return &shipment, nil
}

func (r *shipmentRepository) GetByTrackingNumber(ctx context.Context, trackingNumber string) (*entity.Shipment, error) {
	var shipment entity.Shipment
	err := r.db.WithContext(ctx).Where("tracking_number = ?", trackingNumber).First(&shipment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("shipment not found")
		}
		return nil, err
	}
	return &shipment, nil
}

func (r *shipmentRepository) Update(ctx context.Context, shipment *entity.Shipment) error {
	return r.db.WithContext(ctx).Save(shipment).Error
}

func (r *shipmentRepository) UpdateStatus(ctx context.Context, id uint, status string) error {
	return r.db.WithContext(ctx).Model(&entity.Shipment{}).Where("id = ?", id).Update("status", status).Error
}

func (r *shipmentRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.Shipment{}).Error
}

func (r *shipmentRepository) GetByStatus(ctx context.Context, status string) ([]entity.Shipment, error) {
	var shipments []entity.Shipment
	err := r.db.WithContext(ctx).Where("status = ?", status).Find(&shipments).Error
	return shipments, err
}

func (r *shipmentRepository) GetAll(ctx context.Context, limit, offset int) ([]entity.Shipment, error) {
	var shipments []entity.Shipment
	err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&shipments).Error
	return shipments, err
}