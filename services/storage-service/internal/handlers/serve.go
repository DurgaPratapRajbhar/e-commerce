package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/DurgaPratapRajbhar/e-commerce/pkg/utils"
	"github.com/DurgaPratapRajbhar/e-commerce/storage-service/internal/storage"
	"github.com/gin-gonic/gin"
)

type ServeHandler struct {
	storage *storage.LocalStorage
}

func NewServeHandler(storage *storage.LocalStorage) *ServeHandler {
	return &ServeHandler{
		storage: storage,
	}
}

func (h *ServeHandler) ServeFile(c *gin.Context) {
	// Get file path from URL params
	filePath := c.Param("filepath")
	if filePath == "" {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.ErrInvalidInput, "File path is required", nil, utils.GenerateRequestID()))
		return
	}

	// Check if file exists
	if !h.storage.FileExists(filePath) {
		c.JSON(http.StatusNotFound, utils.ErrorResponse(utils.ErrNotFound, "File not found", nil, utils.GenerateRequestID()))
		return
	}

	// Get file from storage
	fullPath := filepath.Join(h.storage.BasePath, filePath)
	c.File(fullPath)
}

// ServeProductImageFull serves full-sized product images
func (h *ServeHandler) ServeProductImageFull(c *gin.Context) {
	filename := c.Param("filename")
	if filename == "" {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.ErrInvalidInput, "Filename is required", nil, utils.GenerateRequestID()))
		return
	}

	filePath := filepath.Join("products", "images", "full", filename)
	fullPath := filepath.Join(h.storage.BasePath, filePath)
	
	// Check if file exists
	if !h.storage.FileExists(filePath) {
		c.JSON(http.StatusNotFound, utils.ErrorResponse(utils.ErrNotFound, "Product image not found", nil, utils.GenerateRequestID()))
		return
	}

	c.File(fullPath)
}

// ServeProductImageMedium serves medium-sized product images
func (h *ServeHandler) ServeProductImageMedium(c *gin.Context) {
	filename := c.Param("filename")
	if filename == "" {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.ErrInvalidInput, "Filename is required", nil, utils.GenerateRequestID()))
		return
	}

	filePath := filepath.Join("products", "images", "medium", filename)
	fullPath := filepath.Join(h.storage.BasePath, filePath)
	
	// Check if file exists
	if !h.storage.FileExists(filePath) {
		c.JSON(http.StatusNotFound, utils.ErrorResponse(utils.ErrNotFound, "Product image not found", nil, utils.GenerateRequestID()))
		return
	}

	c.File(fullPath)
}

// ServeProductImageThumbnail serves thumbnail-sized product images
func (h *ServeHandler) ServeProductImageThumbnail(c *gin.Context) {
	filename := c.Param("filename")
	if filename == "" {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.ErrInvalidInput, "Filename is required", nil, utils.GenerateRequestID()))
		return
	}

	filePath := filepath.Join("products", "images", "thumbnails", filename)
	fullPath := filepath.Join(h.storage.BasePath, filePath)
	
	// Check if file exists
	if !h.storage.FileExists(filePath) {
		c.JSON(http.StatusNotFound, utils.ErrorResponse(utils.ErrNotFound, "Product image thumbnail not found", nil, utils.GenerateRequestID()))
		return
	}

	c.File(fullPath)
}

// ServeAvatar serves user avatars
func (h *ServeHandler) ServeAvatar(c *gin.Context) {
	filename := c.Param("filename")
	if filename == "" {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.ErrInvalidInput, "Filename is required", nil, utils.GenerateRequestID()))
		return
	}

	filePath := filepath.Join("users", "avatars", filename)
	fullPath := filepath.Join(h.storage.BasePath, filePath)
	
	// Check if file exists
	if !h.storage.FileExists(filePath) {
		c.JSON(http.StatusNotFound, utils.ErrorResponse(utils.ErrNotFound, "Avatar not found", nil, utils.GenerateRequestID()))
		return
	}

	c.File(fullPath)
}

// ServeCategoryBanner serves category banners
func (h *ServeHandler) ServeCategoryBanner(c *gin.Context) {
	filename := c.Param("filename")
	if filename == "" {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.ErrInvalidInput, "Filename is required", nil, utils.GenerateRequestID()))
		return
	}

	filePath := filepath.Join("categories", "banners", filename)
	fullPath := filepath.Join(h.storage.BasePath, filePath)
	
	// Check if file exists
	if !h.storage.FileExists(filePath) {
		c.JSON(http.StatusNotFound, utils.ErrorResponse(utils.ErrNotFound, "Category banner not found", nil, utils.GenerateRequestID()))
		return
	}

	c.File(fullPath)
}

// ServeInvoice serves invoice documents
func (h *ServeHandler) ServeInvoice(c *gin.Context) {
	filename := c.Param("filename")
	if filename == "" {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.ErrInvalidInput, "Filename is required", nil, utils.GenerateRequestID()))
		return
	}

	filePath := filepath.Join("documents", "invoices", filename)
	fullPath := filepath.Join(h.storage.BasePath, filePath)
	
	// Check if file exists
	if !h.storage.FileExists(filePath) {
		c.JSON(http.StatusNotFound, utils.ErrorResponse(utils.ErrNotFound, "Invoice not found", nil, utils.GenerateRequestID()))
		return
	}

	c.File(fullPath)
}

// ServeReceipt serves receipt documents
func (h *ServeHandler) ServeReceipt(c *gin.Context) {
	filename := c.Param("filename")
	if filename == "" {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.ErrInvalidInput, "Filename is required", nil, utils.GenerateRequestID()))
		return
	}

	filePath := filepath.Join("documents", "receipts", filename)
	fullPath := filepath.Join(h.storage.BasePath, filePath)
	
	// Check if file exists
	if !h.storage.FileExists(filePath) {
		c.JSON(http.StatusNotFound, utils.ErrorResponse(utils.ErrNotFound, "Receipt not found", nil, utils.GenerateRequestID()))
		return
	}

	c.File(fullPath)
}

// ListAllFiles lists all files
func (h *ServeHandler) ListAllFiles(c *gin.Context) {
	// This would return a list of all files, but for now we'll return a simple response
	c.JSON(http.StatusOK, utils.SuccessResponse(nil, "File listing not implemented", utils.GenerateRequestID()))
}


// ListFilesByCategory lists files by category
func (h *ServeHandler) ListFilesByCategory(c *gin.Context) {
	category := c.Param("category")
	if category == "" {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.ErrInvalidInput, "Category is required", nil, utils.GenerateRequestID()))
		return
	}

	// This would return a list of files in the category, but for now we'll return a simple response
	c.JSON(http.StatusOK, utils.SuccessResponse(nil, "File listing by category not implemented", utils.GenerateRequestID()))
}