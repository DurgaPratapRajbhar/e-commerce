package repository

import (
	"fmt"
	"log"

	"github.com/DurgaPratapRajbhar/e-commerce/product-service/models"
	"github.com/DurgaPratapRajbhar/e-commerce/product-service/repository"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) repository.ProductRepository {
	return &productRepository{db: db}
}

// func (r *productRepository) CreateProduct(product *models.Product) error {
// 	var existingProduct models.Product
// 	if err := r.db.Where("slug = ?", product.Slug).First(&existingProduct).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
// 		return fmt.Errorf("error checking existing product: %v", err)
// 	} else if err == nil {
// 		return fmt.Errorf("product with slug '%s' already exists", product.Slug)
// 	}

// 	err := r.db.Transaction(func(tx *gorm.DB) error {
// 		if err := tx.Create(product).Error; err != nil {
// 			return fmt.Errorf("failed to create product: %v", err)
// 		}

// 		// Check for SKU uniqueness before creating variants
// 		var skus []string
// 		for _, variant := range product.Variants {
// 			skus = append(skus, variant.SKU)
// 		}

// 		var count int64
// 		if err := tx.Model(&models.ProductVariant{}).Where("sku IN ?", skus).Count(&count).Error; err != nil {
// 			return fmt.Errorf("failed to check SKU uniqueness: %v", err)
// 		}

// 		// fmt.Println(skus, count)
// 		// if count > 0 {
// 		// 	return fmt.Errorf("one or more SKUs already exist")
// 		// }

// 		// Assign ProductID and reset IDs
// 		for i := range product.Variants {
// 			product.Variants[i].ProductID = product.ID
// 			product.Variants[i].ID = 0
// 		}
// 		for i := range product.Attributes {
// 			product.Attributes[i].ProductID = product.ID
// 			product.Attributes[i].ID = 0
// 		}

// 		// Bulk insert for better performance
// 		if len(product.Variants) > 0 {
// 			if err := tx.Create(&product.Variants).Error; err != nil {
// 				return fmt.Errorf("failed to create variants: %v", err)
// 			}
// 		}
// 		if len(product.Attributes) > 0 {
// 			if err := tx.Create(&product.Attributes).Error; err != nil {
// 				return fmt.Errorf("failed to create attributes: %v", err)
// 			}
// 		}

// 		return nil
// 	})

// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (r *productRepository) CreateProduct(product *models.Product) error {
	// return r.db.Create(product).Error

	var existingProduct models.Product
	if err := r.db.Where("slug = ?", product.Slug).First(&existingProduct).Error; err == nil {

		return fmt.Errorf("product with slug '%s' already exists", product.Slug)
	}

	//fmt.Println(product)

	// If no duplicate, proceed with creation
	return r.db.Create(product).Error

}

// func (r *productRepository) GetProduct(id uint) (*models.Product, error) {
// 	var product models.Product
// 	err := r.db.Preload("Category").First(&product, id).Error
// 	return &product, err
// }

// func (r *productRepository) UpdateProduct(id uint, product *models.Product) error {

// 	fmt.Println("product", product, "productproductproduct")
// 	return r.db.Preload("Category").Where("id = ?", id).Updates(&product).Error
// }

// func (r *productRepository) UpdateProduct(id uint, product *models.Product) error {
// 	existingProduct := &models.Product{}

// 	if err := r.db.First(existingProduct, id).Error; err != nil {
// 		return err // Return an error if product is not found
// 	}

// 	fmt.Println(id, "id", product)
// 	return r.db.Model(existingProduct).Updates(product).Error
// }

func (r *productRepository) UpdateProduct(id uint, product *models.Product) error {
	existingProduct := &models.Product{}
	if err := r.db.First(existingProduct, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("product with ID %d not found", id)
		}
		return fmt.Errorf("failed to fetch product: %v", err)
	}

	// Log the incoming product data
	log.Printf("Updating product ID %d with data: %+v", id, product)

	// Update scalar fields, including zero values
	updates := map[string]interface{}{
		"Name":          product.Name,
		"Slug":          product.Slug,
		"Description":   product.Description,
		"Price":         product.Price,
		"Discount":      product.Discount,
		"Stock":         product.Stock,
		"SKU":           product.SKU,
		"Status":        product.Status,
		"Brand":         product.Brand,
		"CategoryID":    product.CategoryID,
		"UoMID":         product.UoMID,
		"QuantityValue": product.QuantityValue, // Handles nil correctly
		"PrimaryImage":  product.PrimaryImage,
		"UpdatedAt":     product.UpdatedAt,
	}

	// Use a transaction for atomicity
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(existingProduct).Updates(updates).Error; err != nil {
			return fmt.Errorf("failed to update product fields: %v", err)
		}

		// Delete and recreate Variants
		if err := tx.Where("product_id = ?", id).Delete(&models.ProductVariant{}).Error; err != nil {
			return fmt.Errorf("failed to delete existing variants: %v", err)
		}
		for i := range product.Variants {
			product.Variants[i].ProductID = uint64(id)
			if err := tx.Create(&product.Variants[i]).Error; err != nil {
				return fmt.Errorf("failed to create variant %d: %v", i, err)
			}
		}

		// Delete and recreate Attributes
		if err := tx.Where("product_id = ?", id).Delete(&models.ProductAttribute{}).Error; err != nil {
			return fmt.Errorf("failed to delete existing attributes: %v", err)
		}
		for i := range product.Attributes {
			product.Attributes[i].ProductID = uint64(id)
			if err := tx.Create(&product.Attributes[i]).Error; err != nil {
				return fmt.Errorf("failed to create attribute %d: %v", i, err)
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *productRepository) DeleteProduct(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}

// func (r *productRepository) GetAllProducts() ([]models.Product, error) {
// 	var products []models.Product
// 	err := r.db.Preload("Category").Find(&products).Error
// 	return products, err
// }

// func (r *productRepository) GetAllProducts(limit, offset int) ([]models.Product, int64, error) {
// 	var products []models.Product
// 	var total int64

// 	// Get total count of products
// 	if err := r.db.Model(&models.Product{}).Count(&total).Error; err != nil {
// 		return nil, 0, err
// 	}

// 	// Apply pagination and preload Category
// 	err := r.db.Preload("Category").Limit(limit).Offset(offset).Find(&products).Error
// 	if err != nil {
// 		return nil, 0, err
// 	}

// 	return products, total, nil
// }

func (r *productRepository) GetAllProducts(limit, offset int) ([]models.Product, int64, error) {
	var products []models.Product
	var total int64

	query := r.db.Model(&models.Product{})
	// if status != "" {
	//     query = query.Where("status = ?", status)
	// }

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Preload("Category").
		Preload("Unit").
		Preload("Attributes").
		Preload("Variants").
		Preload("Variants.Unit").
		Preload("Images").
		Limit(limit).
		Offset(offset).
		Find(&products).Error
	if err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

func (r *productRepository) GetProduct(id uint) (*models.Product, error) {
	var product models.Product

	err := r.db.Preload("Category").
		Preload("Unit").
		Preload("Attributes").
		Preload("Variants").
		Preload("Variants.Unit").
		Preload("Images").
		Where("id = ?", id).
		First(&product).Error

	if err != nil {
		return nil, err
	}

	return &product, nil
}
