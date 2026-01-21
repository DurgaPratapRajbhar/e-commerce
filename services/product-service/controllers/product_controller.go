package controllers

import (
	"crypto/rand"
	"fmt"
	"html"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/DurgaPratapRajbhar/e-commerce/pkg/utils"
	"github.com/DurgaPratapRajbhar/e-commerce/product-service/models"
	"github.com/DurgaPratapRajbhar/e-commerce/product-service/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ProductController struct {
	productService services.ProductService
	validate       *validator.Validate
	logger         *log.Logger
}

func NewProductController(productService services.ProductService) *ProductController {
	validate := validator.New()
	validate.RegisterValidation("regexp", func(fl validator.FieldLevel) bool {
		pattern := fl.Param()
		match, err := regexp.MatchString(pattern, fl.Field().String())
		return err == nil && match
	})

	return &ProductController{
		productService: productService,
		validate:       validate,
		logger:         log.Default(),
	}
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product with variants and attributes
// @Tags products
// @Accept json
// @Produce json
// @Param product body models.Product true "Product data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /products [post]
func (pc *ProductController) CreateProduct(c *gin.Context) {
	var productInput struct {
		models.Product
		Variants   []models.ProductVariant   `json:"variants,omitempty"`
		Attributes []models.ProductAttribute `json:"attributes,omitempty"`
	}

	if err := c.ShouldBindJSON(&productInput); err != nil {
		validationErrors := utils.ParseValidationErrors(err.Error())
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	// Use provided slug and SKU, or generate if empty
	if productInput.Slug == "" {
		productInput.Slug = GenerateSlug(productInput.Name)
	}
	if productInput.SKU == "" {
		productInput.SKU = GenerateSKU(productInput.Name, productInput.Brand)
	}

	if err := productInput.ValidateBasic(pc.validate); err != nil {
		var errorMessages []map[string]string
		for _, e := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, map[string]string{
				"field":   e.Field(),
				"message": fmt.Sprintf("The %s field is %s", e.Field(), e.Tag()),
			})
		}
		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrValidationFailed, "Invalid input data", errorMessages, requestID)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Validate and prepare variants
	for i := range productInput.Variants {
		variant := &productInput.Variants[i]
		if err := pc.validate.Struct(variant); err != nil {
			var errorMessages []map[string]string
			for _, e := range err.(validator.ValidationErrors) {
				errorMessages = append(errorMessages, map[string]string{
					"field":   e.Field(),
					"message": fmt.Sprintf("Variant %d: %s field is %s", i, e.Field(), e.Tag()),
				})
			}
			requestID := utils.GenerateRequestID()
			response := utils.ErrorResponse(utils.ErrValidationFailed, "Invalid variant data", errorMessages, requestID)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		// fmt.Println(productInput.Name)
		variant.ID = 0 // Ensure new variants get auto-incremented IDs
		if variant.SKU == "" {
			variant.SKU = GenerateSKU(fmt.Sprintf("%s-variant-%d", productInput.Name, i), productInput.Brand)
		}
		// fmt.Printf("Variant %d: ID=%d, Name=%s, QuantityValue=%f, SKU=%s", i, variant.ID, variant.Name, variant.QuantityValue, variant.SKU)

	}

	// Validate attributes
	for i := range productInput.Attributes {
		attr := &productInput.Attributes[i]
		if err := pc.validate.Struct(attr); err != nil {
			var errorMessages []map[string]string
			for _, e := range err.(validator.ValidationErrors) {
				errorMessages = append(errorMessages, map[string]string{
					"field":   e.Field(),
					"message": fmt.Sprintf("Attribute %d: %s field is %s", i, e.Field(), e.Tag()),
				})
			}
			requestID := utils.GenerateRequestID()
			response := utils.ErrorResponse(utils.ErrValidationFailed, "Invalid attribute data", errorMessages, requestID)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		attr.ID = 0 // Ensure new attributes get auto-incremented IDs
	}

	newProduct := models.Product{
		Name:          productInput.Name,
		Slug:          productInput.Slug,
		Description:   productInput.Description,
		Price:         productInput.Price,
		Discount:      productInput.Discount,
		Stock:         productInput.Stock,
		SKU:           productInput.SKU,
		Status:        productInput.Status,
		Brand:         productInput.Brand,
		CategoryID:    productInput.CategoryID,
		UoMID:         productInput.UoMID,
		QuantityValue: productInput.QuantityValue,
		PrimaryImage:  productInput.PrimaryImage,
		Variants:      productInput.Variants,
		Attributes:    productInput.Attributes,
	}

	if err := pc.productService.CreateProduct(&newProduct); err != nil {
		// pc.logger.Printf("Failed to create product: %v", err)
		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to create product", err.Error(), requestID)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	requestID := utils.GenerateRequestID()
	c.JSON(http.StatusCreated, utils.SuccessResponse(newProduct, "Product created successfully", requestID))
}

// GetProduct godoc
// @Summary Get a product by ID
// @Description Get a product by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /products/{id} [get]
func (pc *ProductController) GetProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		validationErrors := []utils.ValidationError{
			{Field: "id", Message: "invalid product ID"},
		}
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	product, err := pc.productService.GetProduct(uint(id))
	if err != nil {
		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrNotFound, "Product not found", "Product not found", requestID)
		c.JSON(http.StatusNotFound, response)
		return
	}

	requestID := utils.GenerateRequestID()
	c.JSON(http.StatusOK, utils.SuccessResponse(product, "Product retrieved successfully", requestID))
}

// GetAllProducts godoc
// @Summary Get all products
// @Description Get all products with pagination
// @Tags products
// @Accept json
// @Produce json
// @Param page query int false "Page number" default=1
// @Param limit query int false "Items per page" default=10
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /products [get]
func (pc *ProductController) GetAllProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 { // Add upper limit
		limit = 10
	}

	offset := (page - 1) * limit

	products, total, err := pc.productService.GetAllProducts(limit, offset)
	if err != nil {
		pc.logger.Printf("Failed to fetch products: %v", err)
		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to fetch products", "Failed to fetch products", requestID)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	requestID := utils.GenerateRequestID()
	c.JSON(http.StatusOK, utils.SuccessResponse(map[string]interface{}{
		"data":       products,
		"page":       page,
		"limit":      limit,
		"total":      total,
		"totalPages": int((total + int64(limit) - 1) / int64(limit)), // Fixed pagination calculation
	}, "Products retrieved successfully", requestID))
}

// UpdateProduct godoc
// @Summary Update a product
// @Description Update a product by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body models.Product true "Product data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /products/{id} [put]
func (pc *ProductController) UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		validationErrors := []utils.ValidationError{
			{Field: "id", Message: "invalid product ID"},
		}
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	existingProduct, err := pc.productService.GetProduct(uint(id))
	if err != nil {
		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrNotFound, "Product not found", "Product not found", requestID)
		c.JSON(http.StatusNotFound, response)
		return
	}

	var productInput struct {
		models.Product
		Variants   []models.ProductVariant   `json:"variants,omitempty"`
		Attributes []models.ProductAttribute `json:"attributes,omitempty"`
	}

	if err := c.ShouldBindJSON(&productInput); err != nil {
		validationErrors := utils.ParseValidationErrors(err.Error())
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	if productInput.Name == "" {
		productInput.Name = existingProduct.Name
	}
	if productInput.SKU == "" {
		productInput.SKU = existingProduct.SKU
	} else {
		productInput.SKU = GenerateSKU(productInput.Name, productInput.Brand)
	}
	if productInput.Slug == "" {
		productInput.Slug = existingProduct.Slug
	} else {
		productInput.Slug = GenerateSlug(productInput.Name)
	}

	// Validate basic product fields
	if err := productInput.ValidateBasic(pc.validate); err != nil {
		var errorMessages []map[string]string
		for _, e := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, map[string]string{
				"field":   e.Field(),
				"message": fmt.Sprintf("The %s field is %s", e.Field(), e.Tag()),
			})
		}
		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrValidationFailed, "Invalid input data", errorMessages, requestID)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Validate variants
	for i, variant := range productInput.Variants {

		fmt.Println(variant.QuantityValue, variant.Name)

		if err := pc.validate.Struct(&variant); err != nil {
			var errorMessages []map[string]string
			for _, e := range err.(validator.ValidationErrors) {
				errorMessages = append(errorMessages, map[string]string{
					"field":   e.Field(),
					"message": fmt.Sprintf("Variant %d: %s field is %s", i, e.Field(), e.Tag()),
				})
			}
			requestID := utils.GenerateRequestID()
			response := utils.ErrorResponse(utils.ErrValidationFailed, "Invalid variant data", errorMessages, requestID)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		// Ensure variant has a SKU
		if variant.SKU == "" {
			variant.SKU = GenerateSKU(fmt.Sprintf("%s-variant-%d", productInput.Name, i), productInput.Brand)
		}
	}

	// Validate attributes
	for i, attr := range productInput.Attributes {
		if err := pc.validate.Struct(&attr); err != nil {
			var errorMessages []map[string]string
			for _, e := range err.(validator.ValidationErrors) {
				errorMessages = append(errorMessages, map[string]string{
					"field":   e.Field(),
					"message": fmt.Sprintf("Attribute %d: %s field is %s", i, e.Field(), e.Tag()),
				})
			}
			requestID := utils.GenerateRequestID()
			response := utils.ErrorResponse(utils.ErrValidationFailed, "Invalid attribute data", errorMessages, requestID)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}

	// Prepare updated product
	updatedProduct := models.Product{
		Name:          productInput.Name,
		Slug:          productInput.Slug,
		Description:   productInput.Description,
		Price:         productInput.Price,
		Discount:      productInput.Discount,
		Stock:         productInput.Stock,
		SKU:           productInput.SKU,
		Status:        productInput.Status,
		Brand:         productInput.Brand,
		CategoryID:    productInput.CategoryID,
		UoMID:         productInput.UoMID,
		QuantityValue: productInput.QuantityValue,
		Variants:      productInput.Variants,
		Attributes:    productInput.Attributes,
		CreatedAt:     existingProduct.CreatedAt, // Preserve original creation time
		UpdatedAt:     time.Now(),
	}

	//quantityValue

	// Update the product
	if err := pc.productService.UpdateProduct(uint(id), &updatedProduct); err != nil {
		pc.logger.Printf("Failed to update product %d: %v", id, err)
		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to update product", err.Error(), requestID)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	requestID := utils.GenerateRequestID()
	c.JSON(http.StatusOK, utils.SuccessResponse(updatedProduct, "Product updated successfully", requestID))
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Delete a product by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products/{id} [delete]
func (pc *ProductController) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		validationErrors := []utils.ValidationError{
			{Field: "id", Message: "invalid product ID"},
		}
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	if err := pc.productService.DeleteProduct(uint(id)); err != nil {
		pc.logger.Printf("Failed to delete product %d: %v", id, err)
		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to delete product", "Failed to delete product", requestID)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	requestID := utils.GenerateRequestID()
	c.JSON(http.StatusOK, utils.SuccessResponse(nil, "Product deleted successfully", requestID))
}

func GenerateSlug(name string) string {
	slug := strings.ToLower(strings.TrimSpace(name))
	slug = html.EscapeString(slug)
	reg := regexp.MustCompile(`[^a-z0-9]+`)
	slug = reg.ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-")
	if slug == "" {
		return "product"
	}
	return slug
}

func GenerateSKU(name, brand string) string {
	brandCode := strings.ToUpper(strings.TrimSpace(brand))
	if brandCode == "" {
		brandCode = "GEN"
	}
	if len(brandCode) > 3 {
		brandCode = brandCode[:3]
	}

	nameCode := strings.ToUpper(strings.TrimSpace(name))
	if nameCode == "" {
		nameCode = "PRD"
	}
	if len(nameCode) > 3 {
		nameCode = nameCode[:3]
	}

	const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	suffix := make([]byte, 4)
	if _, err := rand.Read(suffix); err == nil {
		for i := range suffix {
			suffix[i] = alphabet[suffix[i]%byte(len(alphabet))]
		}
	} else {
		// Fallback if crypto/rand fails
		now := time.Now().UnixNano()
		return fmt.Sprintf("%s-%s-%d", brandCode, nameCode, now%10000)
	}

	return fmt.Sprintf("%s-%s-%s", brandCode, nameCode, string(suffix))
}
