package domain

import (
	"context"
	"github.com/DurgaPratapRajbhar/e-commerce/inventory-service/internal/domain/entity"
)

type InventoryRepository interface {
	Create(ctx context.Context, inventory *entity.Inventory) error
	GetByProductAndVariant(ctx context.Context, productID, variantID uint) (*entity.Inventory, error)
	UpdateQuantity(ctx context.Context, productID, variantID uint, quantityChange int) error
	Update(ctx context.Context, inventory *entity.Inventory) error
	Delete(ctx context.Context, productID, variantID uint) error
	GetLowStock(ctx context.Context, threshold int) ([]entity.Inventory, error)
	GetAll(ctx context.Context) ([]entity.Inventory, error)
}

type InventoryTransactionRepository interface {
	Create(ctx context.Context, transaction *entity.InventoryTransaction) error
	GetByProduct(ctx context.Context, productID uint) ([]entity.InventoryTransaction, error)
	GetByProductAndVariant(ctx context.Context, productID, variantID uint) ([]entity.InventoryTransaction, error)
	GetByReferenceID(ctx context.Context, referenceID uint) ([]entity.InventoryTransaction, error)
	GetRecent(ctx context.Context, limit int) ([]entity.InventoryTransaction, error)
}