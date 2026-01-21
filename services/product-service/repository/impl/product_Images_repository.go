package repository

import (
	"fmt"

	"github.com/DurgaPratapRajbhar/e-commerce/pkg/utils"

	"github.com/DurgaPratapRajbhar/e-commerce/product-service/models"
	"github.com/DurgaPratapRajbhar/e-commerce/product-service/repository"

	"gorm.io/gorm"
)

type productImagesRepository struct {
	db *gorm.DB
}

// NewProductImagesRepository creates a new repository instance
func NewProductImagesRepository(db *gorm.DB) repository.ProductImagesRepository {
	return &productImagesRepository{db: db}
}

// CreateProductImage stores metadata in the database
func (r *productImagesRepository) CreateProductImage(productImage *models.ProductImage) error {

	// fmt.Println("productImage", productImage)
	return r.db.Create(productImage).Error
}

// GetProductImage retrieves a single product image by ID
func (r *productImagesRepository) GetProductImage(id uint64) (*models.ProductImage, error) {
	var productImage models.ProductImage
	err := r.db.First(&productImage, id).Error
	return &productImage, err
}

// GetAllProductImages retrieves all images for a product
func (r *productImagesRepository) GetAllProductImages(productID uint64) ([]models.ProductImage, error) {
	var productImages []models.ProductImage
	err := r.db.Where("product_id = ?", productID).Find(&productImages).Error
	return productImages, err
}

// UpdateProductImage updates an image's metadata (e.g., alt text)
func (r *productImagesRepository) UpdateProductImage(id uint64, productImage *models.ProductImage) error {
	return r.db.Model(&models.ProductImage{}).Where("id = ?", id).Updates(productImage).Error
}

// DeleteProductImage removes an image record from the database and deletes the file
func (r *productImagesRepository) DeleteProductImage(id uint64) error {
	var productImage models.ProductImage

	// Fetch image metadata
	if err := r.db.First(&productImage, id).Error; err != nil {
		return err
	}

	// Delete file using shared utility function
	if err := utils.DeleteImage(productImage.ImageURL); err != nil {
		fmt.Printf("Warning: could not delete file %s: %v\n", productImage.ImageURL, err)
	}

	// Remove record from the database
	return r.db.Delete(&productImage).Error
}

// SetPrimaryProductImage marks an image as the primary image for a product
func (r *productImagesRepository) SetPrimaryProductImage(id uint64) error {
	var productImage models.ProductImage

	// Get current image data
	if err := r.db.First(&productImage, id).Error; err != nil {
		return err
	}

	// Remove primary flag from all other images of the same product
	r.db.Model(&models.ProductImage{}).
		Where("product_id = ?", productImage.ProductID).
		Update("is_primary", false)

	// Set primary flag for the selected image
	return r.db.Model(&models.ProductImage{}).Where("id = ?", id).Update("is_primary", true).Error
}

// StoreImage saves an image file in the shared `github.com/DurgaPratapRajbhar/e-commerce/image/` directory
func (r *productImagesRepository) StoreImage(file []byte, filename string) (string, error) {
	// Use shared utility function for storing image
	return utils.StoreImage(file, filename)
}
