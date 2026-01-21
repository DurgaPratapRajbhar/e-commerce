package repository

import (
	"github.com/DurgaPratapRajbhar/e-commerce/product-service/models"
	"github.com/DurgaPratapRajbhar/e-commerce/product-service/repository"

	"gorm.io/gorm"
)

type ProductReviewRepository struct {
	db *gorm.DB
}

func NewProductReviewRepository(db *gorm.DB) repository.ProductReviewRepository {
	return &ProductReviewRepository{db: db}
}

func (r *ProductReviewRepository) CreateReview(review *models.ProductReview) error {
	return r.db.Create(review).Error
}

func (r *ProductReviewRepository) GetReview(id uint) (*models.ProductReview, error) {
	var review models.ProductReview
	if err := r.db.First(&review, id).Error; err != nil {
		return nil, err
	}
	return &review, nil
}

func (r *ProductReviewRepository) UpdateReview(id uint, review *models.ProductReview) error {
	return r.db.Model(&models.ProductReview{}).Where("id = ?", id).Updates(review).Error
}

func (r *ProductReviewRepository) DeleteReview(id uint) error {
	return r.db.Delete(&models.ProductReview{}, id).Error
}

func (r *ProductReviewRepository) GetAllReviews(productID uint) ([]models.ProductReview, error) {
	var reviews []models.ProductReview
	if err := r.db.Where("product_id = ?", productID).Find(&reviews).Error; err != nil {
		return nil, err
	}
	return reviews, nil
}
