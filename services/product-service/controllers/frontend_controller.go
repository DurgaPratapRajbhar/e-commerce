package controllers

import (
	"net/http"

	"github.com/DurgaPratapRajbhar/e-commerce/product-service/services"

	"github.com/gin-gonic/gin"
)

type FrontendController struct {
	service services.FrontendService
}

func NewFrontendController(service services.FrontendService) *FrontendController {
	return &FrontendController{service: service}
}

func (c *FrontendController) GetProductsByCategorySlug(ctx *gin.Context) {
	slug := ctx.Param("slug") // Get slug from URL

	products, err := c.service.GetProductsByCategorySlug(slug)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": products})
}

func (c *FrontendController) GetProductData(ctx *gin.Context) {
	// productID, err := strconv.Atoi(ctx.Param("product_id"))
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
	// 	return
	// }

	// product, err := c.service.GetProductData(uint(productID))
	// if err != nil {
	// 	ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
	// 	return
	// }

	// ctx.JSON(http.StatusOK, product)

	slug := ctx.Param("slug") // Get slug from URL

	products, err := c.service.GetProductData(slug)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": products})
}

func (c *FrontendController) ProductSearch(ctx *gin.Context) {

	search := ctx.Query("query")

	products, err := c.service.ProductSearch(search)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": products})
}
