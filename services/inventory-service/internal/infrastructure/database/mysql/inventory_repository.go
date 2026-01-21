package mysql

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"github.com/DurgaPratapRajbhar/e-commerce/inventory-service/internal/domain/entity"
	"github.com/DurgaPratapRajbhar/e-commerce/inventory-service/internal/domain"
)

type inventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) domain.InventoryRepository {
	return &inventoryRepository{db: db}
}

func (r *inventoryRepository) Create(ctx context.Context, inventory *entity.Inventory) error {
	return r.db.WithContext(ctx).Create(inventory).Error
}

func (r *inventoryRepository) GetByProductAndVariant(ctx context.Context, productID, variantID uint) (*entity.Inventory, error) {
	var inventory entity.Inventory
	err := r.db.WithContext(ctx).Where("product_id = ? AND variant_id = ?", productID, variantID).First(&inventory).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("inventory not found")
		}
		return nil, err
	}
	return &inventory, nil
}

func (r *inventoryRepository) UpdateQuantity(ctx context.Context, productID, variantID uint, quantityChange int) error {
	var inventory entity.Inventory
	err := r.db.WithContext(ctx).Where("product_id = ? AND variant_id = ?", productID, variantID).First(&inventory).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("inventory not found")
		}
		return err
	}

	// Update the quantity
	inventory.Quantity += quantityChange
	if inventory.Quantity < 0 {
		return errors.New("insufficient inventory")
	}

	return r.db.WithContext(ctx).Model(&inventory).Update("quantity", inventory.Quantity).Error
}

func (r *inventoryRepository) Update(ctx context.Context, inventory *entity.Inventory) error {
	return r.db.WithContext(ctx).Save(inventory).Error
}

func (r *inventoryRepository) Delete(ctx context.Context, productID, variantID uint) error {
	return r.db.WithContext(ctx).Where("product_id = ? AND variant_id = ?", productID, variantID).Delete(&entity.Inventory{}).Error
}

func (r *inventoryRepository) GetLowStock(ctx context.Context, threshold int) ([]entity.Inventory, error) {
	var inventories []entity.Inventory
	err := r.db.WithContext(ctx).Where("quantity <= ?", threshold).Find(&inventories).Error
	return inventories, err
}

func (r *inventoryRepository) GetAll(ctx context.Context) ([]entity.Inventory, error) {
	var inventories []entity.Inventory
	err := r.db.WithContext(ctx).Find(&inventories).Error
	return inventories, err
}