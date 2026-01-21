package services

import (
	"github.com/DurgaPratapRajbhar/e-commerce/product-service/models"
	"github.com/DurgaPratapRajbhar/e-commerce/product-service/repository"
	"github.com/DurgaPratapRajbhar/e-commerce/product-service/services"
)

type frontendService struct {
	repo repository.FrontendRepository
}

func NewFrontendService(repo repository.FrontendRepository) services.FrontendService {
	return &frontendService{repo: repo}
}

func (s *frontendService) GetProductData(slug string) ([]models.Product, error) {
	return s.repo.GetProductData(slug)
}

func (s *frontendService) GetProductsByCategorySlug(slug string) ([]models.Product, error) {
	return s.repo.GetProductsByCategorySlug(slug)
}

func (s *frontendService) ProductSearch(search string) ([]models.Product, error) {
	return s.repo.ProductSearch(search)
}
