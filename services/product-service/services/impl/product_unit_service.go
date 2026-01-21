package services

import (
	"github.com/DurgaPratapRajbhar/e-commerce/product-service/models"
	"github.com/DurgaPratapRajbhar/e-commerce/product-service/repository"
)

type ProductUnitServiceImpl struct {
	repo repository.ProductUnitRepository
}

func NewProductUnitService(repo repository.ProductUnitRepository) *ProductUnitServiceImpl {
	return &ProductUnitServiceImpl{repo: repo}
}

func (s *ProductUnitServiceImpl) CreateUnit(Unit *models.UnitOfMeasurement) error {
	// Business logic or validation can be added here
	return s.repo.CreateUnit(Unit)
}

func (s *ProductUnitServiceImpl) GetUnit(id uint) (*models.UnitOfMeasurement, error) {
	return s.repo.GetUnit(id)
}

func (s *ProductUnitServiceImpl) UpdateUnit(id uint, Unit *models.UnitOfMeasurement) error {
	// You can add business logic for updates if needed
	return s.repo.UpdateUnit(id, Unit)
}

func (s *ProductUnitServiceImpl) DeleteUnit(id uint) error {
	return s.repo.DeleteUnit(id)
}

func (s *ProductUnitServiceImpl) GetAllUnit() ([]models.UnitOfMeasurement, error) {
	return s.repo.GetAllUnits()
}
