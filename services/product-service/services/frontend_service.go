package services

import "github.com/DurgaPratapRajbhar/e-commerce/product-service/models"

type FrontendService interface {
	GetProductData(slug string) ([]models.Product, error)
	GetProductsByCategorySlug(slug string) ([]models.Product, error)
	ProductSearch(search string) ([]models.Product, error)
}
