package services

import (
	"github.com/DurgaPratapRajbhar/e-commerce/product-service/models"
	"github.com/DurgaPratapRajbhar/e-commerce/product-service/repository"
	"github.com/DurgaPratapRajbhar/e-commerce/product-service/services"
)

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) services.ProductService {
	return &productService{repo: repo}
}

func (s *productService) CreateProduct(product *models.Product) error {
	return s.repo.CreateProduct(product)
}

func (s *productService) GetProduct(id uint) (*models.Product, error) {
	return s.repo.GetProduct(id)
}

func (s *productService) UpdateProduct(id uint, product *models.Product) error {
	return s.repo.UpdateProduct(id, product)
}

func (s *productService) DeleteProduct(id uint) error {
	return s.repo.DeleteProduct(id)
}

func (s *productService) GetAllProducts(limit, offset int) ([]models.Product, int64, error) {
	return s.repo.GetAllProducts(limit, offset)
}
