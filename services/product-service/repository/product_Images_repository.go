package repository

import "github.com/DurgaPratapRajbhar/e-commerce/product-service/models"

type ProductImagesRepository interface {
	CreateProductImage(productImage *models.ProductImage) error
	GetProductImage(id uint64) (*models.ProductImage, error)
	GetAllProductImages(productID uint64) ([]models.ProductImage, error)
	UpdateProductImage(id uint64, productImage *models.ProductImage) error
	DeleteProductImage(id uint64) error
	SetPrimaryProductImage(id uint64) error
	StoreImage(file []byte, filename string) (string, error)
}
