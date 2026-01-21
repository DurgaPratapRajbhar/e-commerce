package services

import (
	"errors"
	"fmt"

	"github.com/DurgaPratapRajbhar/e-commerce/product-service/models"
	"github.com/DurgaPratapRajbhar/e-commerce/product-service/repository"

	"github.com/DurgaPratapRajbhar/e-commerce/pkg/utils"

	"github.com/DurgaPratapRajbhar/e-commerce/product-service/services"
)

// ProductImagesService implementation
type productImagesService struct {
	repo repository.ProductImagesRepository
}

// NewProductImagesService creates a new service instance
func NewProductImagesService(repo repository.ProductImagesRepository) services.ProductImagesServices {
	return &productImagesService{repo: repo}
}

// CreateProductImage saves metadata in the database
func (s *productImagesService) CreateProductImage(productImage *models.ProductImage) error {
	return s.repo.CreateProductImage(productImage)
}

// GetProductImage retrieves a single product image by ID
func (s *productImagesService) GetProductImage(id uint64) (*models.ProductImage, error) {
	return s.repo.GetProductImage(id)
}

// GetAllProductImages retrieves all images for a product
func (s *productImagesService) GetAllProductImages(productID uint64) ([]models.ProductImage, error) {
	return s.repo.GetAllProductImages(productID)
}

// UpdateProductImage updates image metadata
func (s *productImagesService) UpdateProductImage(id uint64, productImage *models.ProductImage) error {
	return s.repo.UpdateProductImage(id, productImage)
}

// DeleteProductImage removes an image from storage and the database
func (s *productImagesService) DeleteProductImage(id uint64) error {
	return s.repo.DeleteProductImage(id)
}

// SetPrimaryProductImage marks an image as the primary one
func (s *productImagesService) SetPrimaryProductImage(id uint64) error {
	return s.repo.SetPrimaryProductImage(id)
}

// ValidateAndStoreImage validates an image file and saves it
func (s *productImagesService) ValidateAndStoreImage(file []byte, filename string) (string, error) {
	// Validate file type
	if !utils.IsValidImageFormat(filename) {
		return "", errors.New("invalid file format: only JPG, PNG, and GIF are allowed")
	}

	// Validate file size (limit: 5MB)
	if len(file) > 5*1024*1024 {
		return "", errors.New("file size exceeds 5MB limit")
	}

	// Store the image using the repository method
	storedPath, err := s.repo.StoreImage(file, filename)
	if err != nil {
		return "", fmt.Errorf("failed to store image: %w", err)
	}

	return storedPath, nil
}
