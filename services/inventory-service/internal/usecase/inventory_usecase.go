package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/DurgaPratapRajbhar/e-commerce/inventory-service/internal/domain"
	"github.com/DurgaPratapRajbhar/e-commerce/inventory-service/internal/domain/entity"
)

type InventoryUseCase struct {
	inventoryRepo        domain.InventoryRepository
	transactionRepo domain.InventoryTransactionRepository
}

func NewInventoryUseCase(
	inventoryRepo domain.InventoryRepository,
	transactionRepo domain.InventoryTransactionRepository,
) *InventoryUseCase {
	return &InventoryUseCase{
		inventoryRepo:        inventoryRepo,
		transactionRepo: transactionRepo,
	}
}

func (uc *InventoryUseCase) CreateInventory(ctx context.Context, inventory *entity.Inventory) error {
	// Check if inventory already exists
	existing, err := uc.inventoryRepo.GetByProductAndVariant(ctx, inventory.ProductID, inventory.VariantID)
	if err == nil && existing != nil {
		return errors.New("inventory already exists for this product and variant")
	}

	return uc.inventoryRepo.Create(ctx, inventory)
}

func (uc *InventoryUseCase) DeleteInventory(ctx context.Context, productID, variantID uint) error {
	return uc.inventoryRepo.Delete(ctx, productID, variantID)
}

func (uc *InventoryUseCase) GetInventory(ctx context.Context, productID, variantID uint) (*entity.Inventory, error) {
	return uc.inventoryRepo.GetByProductAndVariant(ctx, productID, variantID)
}

func (uc *InventoryUseCase) UpdateInventory(ctx context.Context, req *entity.InventoryUpdateRequest) error {
	// Create transaction record
	transaction := &entity.InventoryTransaction{
		ProductID:       req.ProductID,
		VariantID:       req.VariantID,
		TransactionType: req.TransactionType,
		Quantity:        req.QuantityChange,
		ReferenceID:     req.ReferenceID,
	}

	// Update inventory quantity
	err := uc.inventoryRepo.UpdateQuantity(ctx, req.ProductID, req.VariantID, req.QuantityChange)
	if err != nil {
		return err
	}

	// Update warehouse location if provided
	if req.WarehouseLocation != nil {
		inventory, err := uc.inventoryRepo.GetByProductAndVariant(ctx, req.ProductID, req.VariantID)
		if err != nil {
			return err
		}
		inventory.WarehouseLocation = *req.WarehouseLocation
		err = uc.inventoryRepo.Update(ctx, inventory)
		if err != nil {
			return err
		}
	}

	// Save transaction
	err = uc.transactionRepo.Create(ctx, transaction)
	if err != nil {
		return fmt.Errorf("failed to record inventory transaction: %w", err)
	}

	return nil
}

func (uc *InventoryUseCase) GetInventoryTransactions(ctx context.Context, productID, variantID uint) ([]entity.InventoryTransaction, error) {
	if productID != 0 && variantID != 0 {
		return uc.transactionRepo.GetByProductAndVariant(ctx, productID, variantID)
	} else if productID != 0 {
		return uc.transactionRepo.GetByProduct(ctx, productID)
	}
	return nil, errors.New("either productID or both productID and variantID must be provided")
}

func (uc *InventoryUseCase) GetLowStockItems(ctx context.Context, threshold int) ([]entity.Inventory, error) {
	return uc.inventoryRepo.GetLowStock(ctx, threshold)
}

func (uc *InventoryUseCase) ReserveInventory(ctx context.Context, productID, variantID uint, quantity int, referenceID uint) error {
	req := &entity.InventoryUpdateRequest{
		ProductID:       productID,
		VariantID:       variantID,
		QuantityChange:  -quantity, // Negative to reduce available quantity
		TransactionType: "reserved",
		ReferenceID:     &referenceID,
	}
	
	return uc.UpdateInventory(ctx, req)
}

func (uc *InventoryUseCase) ReleaseReservedInventory(ctx context.Context, productID, variantID uint, quantity int, referenceID uint) error {
	req := &entity.InventoryUpdateRequest{
		ProductID:       productID,
		VariantID:       variantID,
		QuantityChange:  quantity, // Positive to increase available quantity
		TransactionType: "in",
		ReferenceID:     &referenceID,
	}
	
	return uc.UpdateInventory(ctx, req)
}

func (uc *InventoryUseCase) GetRecentTransactions(ctx context.Context, limit int) ([]entity.InventoryTransaction, error) {
	return uc.transactionRepo.GetRecent(ctx, limit)
}

func (uc *InventoryUseCase) GetInventoryList(ctx context.Context) ([]entity.Inventory, error) {
	return uc.inventoryRepo.GetAll(ctx)
}