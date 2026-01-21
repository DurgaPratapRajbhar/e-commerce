package mysql

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"github.com/DurgaPratapRajbhar/e-commerce/payment-service/internal/domain/entity"
	"github.com/DurgaPratapRajbhar/e-commerce/payment-service/internal/domain"
)

type refundRepository struct {
	db *gorm.DB
}

func NewRefundRepository(db *gorm.DB) domain.RefundRepository {
	return &refundRepository{db: db}
}

func (r *refundRepository) Create(ctx context.Context, refund *entity.Refund) error {
	return r.db.WithContext(ctx).Create(refund).Error
}

func (r *refundRepository) GetByID(ctx context.Context, id uint) (*entity.Refund, error) {
	var refund entity.Refund
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&refund).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("refund not found")
		}
		return nil, err
	}
	return &refund, nil
}

func (r *refundRepository) GetByPaymentID(ctx context.Context, paymentID uint) ([]entity.Refund, error) {
	var refunds []entity.Refund
	err := r.db.WithContext(ctx).Where("payment_id = ?", paymentID).Find(&refunds).Error
	return refunds, err
}

func (r *refundRepository) GetByOrderID(ctx context.Context, orderID uint) ([]entity.Refund, error) {
	var refunds []entity.Refund
	err := r.db.WithContext(ctx).Where("order_id = ?", orderID).Find(&refunds).Error
	return refunds, err
}

func (r *refundRepository) GetByStatus(ctx context.Context, status string) ([]entity.Refund, error) {
	var refunds []entity.Refund
	err := r.db.WithContext(ctx).Where("status = ?", status).Find(&refunds).Error
	return refunds, err
}

func (r *refundRepository) Update(ctx context.Context, refund *entity.Refund) error {
	return r.db.WithContext(ctx).Save(refund).Error
}

func (r *refundRepository) UpdateStatus(ctx context.Context, id uint, status string) error {
	return r.db.WithContext(ctx).Model(&entity.Refund{}).Where("id = ?", id).Update("status", status).Error
}

func (r *refundRepository) GetAll(ctx context.Context, limit, offset int) ([]entity.Refund, error) {
	var refunds []entity.Refund
	err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&refunds).Error
	return refunds, err
}