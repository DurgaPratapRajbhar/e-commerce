package domain

import (
	"context"
	"github.com/DurgaPratapRajbhar/e-commerce/payment-service/internal/domain/entity"
)

type PaymentRepository interface {
	Create(ctx context.Context, payment *entity.Payment) error
	GetByID(ctx context.Context, id uint) (*entity.Payment, error)
	GetByOrderID(ctx context.Context, orderID uint) (*entity.Payment, error)
	GetByUserID(ctx context.Context, userID uint) ([]entity.Payment, error)
	GetByTransactionID(ctx context.Context, transactionID string) (*entity.Payment, error)
	GetByStatus(ctx context.Context, status string) ([]entity.Payment, error)
	Update(ctx context.Context, payment *entity.Payment) error
	UpdateStatus(ctx context.Context, id uint, status string) error
	Delete(ctx context.Context, id uint) error
	GetAll(ctx context.Context, limit, offset int) ([]entity.Payment, error)
}

type RefundRepository interface {
	Create(ctx context.Context, refund *entity.Refund) error
	GetByID(ctx context.Context, id uint) (*entity.Refund, error)
	GetByPaymentID(ctx context.Context, paymentID uint) ([]entity.Refund, error)
	GetByOrderID(ctx context.Context, orderID uint) ([]entity.Refund, error)
	GetByStatus(ctx context.Context, status string) ([]entity.Refund, error)
	Update(ctx context.Context, refund *entity.Refund) error
	UpdateStatus(ctx context.Context, id uint, status string) error
	GetAll(ctx context.Context, limit, offset int) ([]entity.Refund, error)
}