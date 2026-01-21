package repository

import (
	"github.com/DurgaPratapRajbhar/e-commerce/product-service/models"
	"github.com/DurgaPratapRajbhar/e-commerce/product-service/repository"

	"gorm.io/gorm"
)

type frontendRepository struct {
	db *gorm.DB
}

func NewFrontendRepository(db *gorm.DB) repository.FrontendRepository {
	return &frontendRepository{db: db}
}

func (r *frontendRepository) GetProductsByCategorySlug(slug string) ([]models.Product, error) {
	var products []models.Product
	var categoryID uint
	var childCategoryIDs []uint

	// Parent ya child category ka ID fetch karna
	err := r.db.Table("categories").Select("id").Where("slug = ?", slug).Scan(&categoryID).Error
	if err != nil {
		return nil, err
	}

	// Check karna ki yeh parent category hai ya nahi
	err = r.db.Table("categories").Select("id").Where("parent_id = ?", categoryID).Pluck("id", &childCategoryIDs).Error
	if err != nil {
		return nil, err
	}

	// Agar child categories mili hain to parent ka ID bhi list me daal do
	if len(childCategoryIDs) > 0 {
		childCategoryIDs = append(childCategoryIDs, categoryID)
	} else {
		// Agar child nahi mili, to sirf ek category ka ID use karna hai
		childCategoryIDs = []uint{categoryID}
	}

	// Products fetch karna
	// err = r.db.
	// 	Table("products").
	// 	Select("products.*, (SELECT image_url FROM product_images WHERE product_images.product_id = products.id AND product_images.is_primary = 1 LIMIT 1) AS primary_image").
	// 	Where("products.category_id IN (?) AND products.status = ?", childCategoryIDs, "active").
	// 	Find(&products).Error

	err = r.db.
		Table("products").
		Select("products.*, product_images.image_url AS primary_image").
		Joins("LEFT JOIN product_images ON product_images.product_id = products.id AND product_images.is_primary = 1").
		Preload("Variants"). // Preload the product variants
		Where("products.category_id IN (?) AND products.status = ?", childCategoryIDs, "active").
		Find(&products).Error

	if err != nil {
		return nil, err
	}

	return products, nil
}

// func (r *frontendRepository) GetProductData(slug string) ([]models.Product, error) {
// 	var products []models.Product

// 	err := r.db.
// 		Preload("Images"). // ✅ Loads all images for each product
// 		Joins("JOIN categories ON categories.id = products.category_id").
// 		Where("products.slug = ? AND products.status = ?", slug, "active").
// 		Find(&products).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	return products, nil
// }

func (r *frontendRepository) GetProductData(slug string) ([]models.Product, error) {
	var products []models.Product

	err := r.db.
		Preload("Images").     // ✅ Loads all product images
		Preload("Category").   // ✅ Loads category details
		Preload("Unit").       // ✅ Loads unit of measurement details
		Preload("Attributes"). // ✅ Loads product attributes
		Preload("Variants").   // ✅ Loads product variants
		Joins("JOIN categories ON categories.id = products.category_id").
		Where("products.slug = ? AND products.status = ?", slug, "active").
		Find(&products).Error

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *frontendRepository) ProductSearch(search string) ([]models.Product, error) {
	var products []models.Product

	query := r.db.
		Preload("Images").
		Preload("Category").
		Joins("LEFT JOIN categories ON categories.id = products.category_id").
		Where("products.status = ?", "active")

	if search != "" {
		query = query.Where("products.name LIKE ? OR products.description LIKE ? OR categories.name LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}
