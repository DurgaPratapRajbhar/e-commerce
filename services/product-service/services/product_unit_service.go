package services

import "github.com/DurgaPratapRajbhar/e-commerce/product-service/models"

type ProductUnitServices interface {
	CreateUnit(Unit *models.UnitOfMeasurement) error
	GetUnit(id uint) (*models.UnitOfMeasurement, error)
	UpdateUnit(id uint, Unit *models.UnitOfMeasurement) error
	DeleteUnit(id uint) error
	GetAllUnit() ([]models.UnitOfMeasurement, error)
}
