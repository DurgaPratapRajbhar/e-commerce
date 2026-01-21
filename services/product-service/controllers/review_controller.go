package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/DurgaPratapRajbhar/e-commerce/product-service/models"
	"github.com/DurgaPratapRajbhar/e-commerce/product-service/services"
)

type ProductReviewController struct {
	productReviewService services.ProductReviewService
	validator            *validator.Validate
}

func NewProductReviewController(productReviewService services.ProductReviewService) *ProductReviewController {
	return &ProductReviewController{
		productReviewService: productReviewService,
		validator:            validator.New(),
	}
}

// CreateReview creates a new product review
// @Summary Create a new product review
// @Description Create a new product review
// @Tags Reviews
// @Accept json
// @Produce json
// @Param review body models.ProductReview true "Review data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /product-reviews [post]
func (prc *ProductReviewController) CreateReview(c *gin.Context) {
	var review models.ProductReview
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the struct
	if err := prc.validator.Struct(&review); err != nil {

		fmt.Println(err)

		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return
	}

	if err := prc.productReviewService.CreateReview(&review); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product review"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product review created successfully", "data": review})
}

// GetReview retrieves a product review by ID
// @Summary Get a review by ID
// @Description Get a review by its ID
// @Tags Reviews
// @Accept json
// @Produce json
// @Param id path int true "Review ID"
// @Success 200 {object} models.ProductReview
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /product-reviews/{id} [get]
func (prc *ProductReviewController) GetReview(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review ID"})
		return
	}

	review, err := prc.productReviewService.GetReview(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product review not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": review})
}

// GetAllReviews retrieves all reviews for a product
// @Summary Get all reviews for a product
// @Description Get all reviews for a product
// @Tags Reviews
// @Accept json
// @Produce json
// @Param product_id query int false "Product ID"
// @Success 200 {array} models.ProductReview
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /product-reviews [get]
func (prc *ProductReviewController) GetAllReviews(c *gin.Context) {
	productID, err := strconv.Atoi(c.DefaultQuery("product_id", "0"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	reviews, err := prc.productReviewService.GetAllReviews(uint(productID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product reviews"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reviews})
}

// UpdateReview updates an existing product review
// @Summary Update a product review
// @Description Update a product review by its ID
// @Tags Reviews
// @Accept json
// @Produce json
// @Param id path int true "Review ID"
// @Param review body models.ProductReview true "Review data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /product-reviews/{id} [put]
func (prc *ProductReviewController) UpdateReview(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review ID"})
		return
	}

	var review models.ProductReview
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the struct
	if err := prc.validator.Struct(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return
	}

	if err := prc.productReviewService.UpdateReview(uint(id), &review); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product review"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product review updated successfully"})
}

// DeleteReview deletes a product review by ID
// @Summary Delete a product review
// @Description Delete a product review by its ID
// @Tags Reviews
// @Accept json
// @Produce json
// @Param id path int true "Review ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /product-reviews/{id} [delete]
func (prc *ProductReviewController) DeleteReview(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review ID"})
		return
	}

	if err := prc.productReviewService.DeleteReview(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product review"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product review deleted successfully"})
}
