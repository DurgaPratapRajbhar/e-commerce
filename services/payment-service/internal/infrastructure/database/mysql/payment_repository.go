package mysql

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"github.com/DurgaPratapRajbhar/e-commerce/payment-service/internal/domain/entity"
	"github.com/DurgaPratapRajbhar/e-commerce/payment-service/internal/domain"
)

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) domain.PaymentRepository {
	return &paymentRepository{db: db}
}

func (r *paymentRepository) Create(ctx context.Context, payment *entity.Payment) error {
	return r.db.WithContext(ctx).Create(payment).Error
}

func (r *paymentRepository) GetByID(ctx context.Context, id uint) (*entity.Payment, error) {
	var payment entity.Payment
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&payment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("payment not found")
		}
		return nil, err
	}
	return &payment, nil
}

func (r *paymentRepository) GetByOrderID(ctx context.Context, orderID uint) (*entity.Payment, error) {
	var payment entity.Payment
	err := r.db.WithContext(ctx).Where("order_id = ?", orderID).First(&payment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("payment not found")
		}
		return nil, err
	}
	return &payment, nil
}

func (r *paymentRepository) GetByUserID(ctx context.Context, userID uint) ([]entity.Payment, error) {
	var payments []entity.Payment
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&payments).Error
	return payments, err
}

func (r *paymentRepository) GetByTransactionID(ctx context.Context, transactionID string) (*entity.Payment, error) {
	var payment entity.Payment
	err := r.db.WithContext(ctx).Where("transaction_id = ?", transactionID).First(&payment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("payment not found")
		}
		return nil, err
	}
	return &payment, nil
}

func (r *paymentRepository) GetByStatus(ctx context.Context, status string) ([]entity.Payment, error) {
	var payments []entity.Payment
	err := r.db.WithContext(ctx).Where("payment_status = ?", status).Find(&payments).Error
	return payments, err
}

func (r *paymentRepository) Update(ctx context.Context, payment *entity.Payment) error {
	return r.db.WithContext(ctx).Save(payment).Error
}

func (r *paymentRepository) UpdateStatus(ctx context.Context, id uint, status string) error {
	return r.db.WithContext(ctx).Model(&entity.Payment{}).Where("id = ?", id).Update("payment_status", status).Error
}

func (r *paymentRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.Payment{}).Error
}

func (r *paymentRepository) GetAll(ctx context.Context, limit, offset int) ([]entity.Payment, error) {
	var payments []entity.Payment
	err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&payments).Error
	return payments, err
}