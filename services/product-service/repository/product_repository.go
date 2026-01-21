package repository

import "github.com/DurgaPratapRajbhar/e-commerce/product-service/models"

type ProductRepository interface {
	CreateProduct(product *models.Product) error
	GetProduct(id uint) (*models.Product, error)
	UpdateProduct(id uint, product *models.Product) error
	DeleteProduct(id uint) error
	GetAllProducts(limit, offset int) ([]models.Product, int64, error)
}
