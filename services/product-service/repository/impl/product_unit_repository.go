package repository

import (
	"github.com/DurgaPratapRajbhar/e-commerce/product-service/models"

	"gorm.io/gorm"
)

type ProductUnitRepositoryImpl struct {
	db *gorm.DB
}

func NewProductUnitRepository(db *gorm.DB) *ProductUnitRepositoryImpl {
	return &ProductUnitRepositoryImpl{db: db}
}

func (r *ProductUnitRepositoryImpl) CreateUnit(Unit *models.UnitOfMeasurement) error {
	return r.db.Create(Unit).Error
}

func (r *ProductUnitRepositoryImpl) GetUnit(id uint) (*models.UnitOfMeasurement, error) {
	var Unit models.UnitOfMeasurement
	if err := r.db.First(&Unit, id).Error; err != nil {
		return nil, err
	}
	return &Unit, nil
}

func (r *ProductUnitRepositoryImpl) UpdateUnit(id uint, Unit *models.UnitOfMeasurement) error {
	return r.db.Model(&models.UnitOfMeasurement{}).Where("id = ?", id).Updates(Unit).Error
}

func (r *ProductUnitRepositoryImpl) DeleteUnit(id uint) error {
	return r.db.Delete(&models.UnitOfMeasurement{}, id).Error
}

func (r *ProductUnitRepositoryImpl) GetAllUnits() ([]models.UnitOfMeasurement, error) {
	var units []models.UnitOfMeasurement
	if err := r.db.Find(&units).Error; err != nil {
		return nil, err
	}
	return units, nil
}
