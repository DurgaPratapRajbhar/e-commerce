package mysql

import (
	"context"
	"gorm.io/gorm"
	"github.com/DurgaPratapRajbhar/e-commerce/inventory-service/internal/domain/entity"
	"github.com/DurgaPratapRajbhar/e-commerce/inventory-service/internal/domain"
)

type inventoryTransactionRepository struct {
	db *gorm.DB
}

func NewInventoryTransactionRepository(db *gorm.DB) domain.InventoryTransactionRepository {
	return &inventoryTransactionRepository{db: db}
}

func (r *inventoryTransactionRepository) Create(ctx context.Context, transaction *entity.InventoryTransaction) error {
	return r.db.WithContext(ctx).Create(transaction).Error
}

func (r *inventoryTransactionRepository) GetByProduct(ctx context.Context, productID uint) ([]entity.InventoryTransaction, error) {
	var transactions []entity.InventoryTransaction
	err := r.db.WithContext(ctx).Where("product_id = ?", productID).Find(&transactions).Error
	return transactions, err
}

func (r *inventoryTransactionRepository) GetByProductAndVariant(ctx context.Context, productID, variantID uint) ([]entity.InventoryTransaction, error) {
	var transactions []entity.InventoryTransaction
	err := r.db.WithContext(ctx).Where("product_id = ? AND variant_id = ?", productID, variantID).Find(&transactions).Error
	return transactions, err
}

func (r *inventoryTransactionRepository) GetByReferenceID(ctx context.Context, referenceID uint) ([]entity.InventoryTransaction, error) {
	var transactions []entity.InventoryTransaction
	err := r.db.WithContext(ctx).Where("reference_id = ?", referenceID).Find(&transactions).Error
	return transactions, err
}

func (r *inventoryTransactionRepository) GetRecent(ctx context.Context, limit int) ([]entity.InventoryTransaction, error) {
	var transactions []entity.InventoryTransaction
	err := r.db.WithContext(ctx).Order("created_at DESC").Limit(limit).Find(&transactions).Error
	return transactions, err
}