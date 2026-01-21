package repository

import "github.com/DurgaPratapRajbhar/e-commerce/product-service/models"

type ProductUnitRepository interface {
	CreateUnit(Unit *models.UnitOfMeasurement) error
	GetUnit(id uint) (*models.UnitOfMeasurement, error)
	UpdateUnit(id uint, Unit *models.UnitOfMeasurement) error
	DeleteUnit(id uint) error
	GetAllUnits() ([]models.UnitOfMeasurement, error)
}
