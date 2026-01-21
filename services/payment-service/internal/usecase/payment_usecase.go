package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"
	"github.com/DurgaPratapRajbhar/e-commerce/payment-service/internal/domain"
	"github.com/DurgaPratapRajbhar/e-commerce/payment-service/internal/domain/entity"
)

type PaymentUseCase struct {
	paymentRepo domain.PaymentRepository
	refundRepo  domain.RefundRepository
}

func NewPaymentUseCase(
	paymentRepo domain.PaymentRepository,
	refundRepo domain.RefundRepository,
) *PaymentUseCase {
	return &PaymentUseCase{
		paymentRepo: paymentRepo,
		refundRepo:  refundRepo,
	}
}

func (uc *PaymentUseCase) CreatePayment(ctx context.Context, req *entity.PaymentCreateRequest) (*entity.Payment, error) {
	// Check if payment already exists for this order
	existing, err := uc.paymentRepo.GetByOrderID(ctx, req.OrderID)
	if err == nil && existing != nil {
		return nil, errors.New("payment already exists for this order")
	}

	// Generate a transaction ID if not provided
	transactionID := fmt.Sprintf("TXN-%d", time.Now().Unix())

	payment := &entity.Payment{
		OrderID:       req.OrderID,
		UserID:        req.UserID,
		Amount:        req.Amount,
		PaymentMethod: req.PaymentMethod,
		PaymentStatus: "pending",
		TransactionID: transactionID,
	}

	err = uc.paymentRepo.Create(ctx, payment)
	if err != nil {
		return nil, fmt.Errorf("failed to create payment: %w", err)
	}

	return payment, nil
}

func (uc *PaymentUseCase) GetPayment(ctx context.Context, id uint) (*entity.Payment, error) {
	return uc.paymentRepo.GetByID(ctx, id)
}

func (uc *PaymentUseCase) GetPaymentByOrderID(ctx context.Context, orderID uint) (*entity.Payment, error) {
	return uc.paymentRepo.GetByOrderID(ctx, orderID)
}

func (uc *PaymentUseCase) GetPaymentByTransactionID(ctx context.Context, transactionID string) (*entity.Payment, error) {
	return uc.paymentRepo.GetByTransactionID(ctx, transactionID)
}

func (uc *PaymentUseCase) GetPaymentsByUserID(ctx context.Context, userID uint) ([]entity.Payment, error) {
	return uc.paymentRepo.GetByUserID(ctx, userID)
}

func (uc *PaymentUseCase) GetPaymentsByStatus(ctx context.Context, status string) ([]entity.Payment, error) {
	return uc.paymentRepo.GetByStatus(ctx, status)
}

func (uc *PaymentUseCase) UpdatePayment(ctx context.Context, id uint, req *entity.PaymentUpdateRequest) (*entity.Payment, error) {
	payment, err := uc.paymentRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided in request
	if req.PaymentStatus != nil {
		payment.PaymentStatus = *req.PaymentStatus
	}
	if req.TransactionID != nil {
		payment.TransactionID = *req.TransactionID
	}
	if req.GatewayResponse != nil {
		payment.GatewayResponse = req.GatewayResponse
	}

	err = uc.paymentRepo.Update(ctx, payment)
	if err != nil {
		return nil, fmt.Errorf("failed to update payment: %w", err)
	}

	return payment, nil
}

func (uc *PaymentUseCase) UpdatePaymentStatus(ctx context.Context, id uint, status string) error {
	// Validate status
	validStatuses := map[string]bool{
		"pending":   true,
		"success":   true,
		"failed":    true,
		"refunded":  true,
	}
	if !validStatuses[status] {
		return errors.New("invalid payment status")
	}

	err := uc.paymentRepo.UpdateStatus(ctx, id, status)
	if err != nil {
		return fmt.Errorf("failed to update payment status: %w", err)
	}

	return nil
}

func (uc *PaymentUseCase) DeletePayment(ctx context.Context, id uint) error {
	return uc.paymentRepo.Delete(ctx, id)
}

func (uc *PaymentUseCase) GetPayments(ctx context.Context, limit, offset int) ([]entity.Payment, error) {
	return uc.paymentRepo.GetAll(ctx, limit, offset)
}

func (uc *PaymentUseCase) CreateRefund(ctx context.Context, req *entity.RefundCreateRequest) (*entity.Refund, error) {
	// Check if payment exists and is eligible for refund
	payment, err := uc.paymentRepo.GetByID(ctx, req.PaymentID)
	if err != nil {
		return nil, fmt.Errorf("payment not found: %w", err)
	}

	if payment.PaymentStatus != "success" && payment.PaymentStatus != "refunded" {
		return nil, errors.New("payment must be successful to create a refund")
	}

	// Check if the refund amount is valid
	if req.Amount > payment.Amount {
		return nil, errors.New("refund amount exceeds payment amount")
	}

	refund := &entity.Refund{
		PaymentID: req.PaymentID,
		OrderID:   req.OrderID,
		Amount:    req.Amount,
		Reason:    req.Reason,
		Status:    "pending",
	}

	err = uc.refundRepo.Create(ctx, refund)
	if err != nil {
		return nil, fmt.Errorf("failed to create refund: %w", err)
	}

	// Update payment status to refunded if it was successful
	if payment.PaymentStatus == "success" {
		payment.PaymentStatus = "refunded"
		err = uc.paymentRepo.Update(ctx, payment)
		if err != nil {
			return nil, fmt.Errorf("failed to update payment status to refunded: %w", err)
		}
	}

	return refund, nil
}

func (uc *PaymentUseCase) GetRefund(ctx context.Context, id uint) (*entity.Refund, error) {
	return uc.refundRepo.GetByID(ctx, id)
}

func (uc *PaymentUseCase) GetRefundsByPaymentID(ctx context.Context, paymentID uint) ([]entity.Refund, error) {
	return uc.refundRepo.GetByPaymentID(ctx, paymentID)
}

func (uc *PaymentUseCase) GetRefundsByOrderID(ctx context.Context, orderID uint) ([]entity.Refund, error) {
	return uc.refundRepo.GetByOrderID(ctx, orderID)
}

func (uc *PaymentUseCase) GetRefundsByStatus(ctx context.Context, status string) ([]entity.Refund, error) {
	return uc.refundRepo.GetByStatus(ctx, status)
}

func (uc *PaymentUseCase) GetRefunds(ctx context.Context, limit, offset int) ([]entity.Refund, error) {
	return uc.refundRepo.GetAll(ctx, limit, offset)
}