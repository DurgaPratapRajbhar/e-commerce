package services

import (
	"fmt"

	"github.com/DurgaPratapRajbhar/e-commerce/product-service/models"
	"github.com/DurgaPratapRajbhar/e-commerce/product-service/repository"
	"github.com/DurgaPratapRajbhar/e-commerce/product-service/services"
)

type ProductReviewService struct {
	repo repository.ProductReviewRepository
}

func NewProductReviewService(repo repository.ProductReviewRepository) services.ProductReviewService {
	return &ProductReviewService{repo: repo}
}

func (s *ProductReviewService) CreateReview(review *models.ProductReview) error {
	if review.Rating < 1 || review.Rating > 5 {
		return fmt.Errorf("rating must be between 1 and 5")
	}

	return s.repo.CreateReview(review)
}

func (s *ProductReviewService) GetReview(id uint) (*models.ProductReview, error) {
	return s.repo.GetReview(id)
}

func (s *ProductReviewService) UpdateReview(id uint, review *models.ProductReview) error {
	return s.repo.UpdateReview(id, review)
}

func (s *ProductReviewService) DeleteReview(id uint) error {
	return s.repo.DeleteReview(id)
}

func (s *ProductReviewService) GetAllReviews(productID uint) ([]models.ProductReview, error) {
	return s.repo.GetAllReviews(productID)
}
