package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/DurgaPratapRajbhar/e-commerce/pkg/utils"
	"github.com/DurgaPratapRajbhar/e-commerce/storage-service/internal/models"
	"github.com/DurgaPratapRajbhar/e-commerce/storage-service/internal/processor"
	"github.com/DurgaPratapRajbhar/e-commerce/storage-service/internal/storage"
	"github.com/DurgaPratapRajbhar/e-commerce/storage-service/internal/validator"
	"github.com/gin-gonic/gin"
)

type UploadHandler struct {
	storage   *storage.LocalStorage
	validator *validator.FileValidator
	processor *processor.ImageProcessor
}

func NewUploadHandler(storage *storage.LocalStorage, validator *validator.FileValidator, processor *processor.ImageProcessor) *UploadHandler {
	return &UploadHandler{
		storage:   storage,
		validator: validator,
		processor: processor,
	}
}

// UploadProductImage handles product image uploads
func (h *UploadHandler) UploadProductImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.ErrInvalidInput, "File is required", nil, utils.GenerateRequestID()))
		return
	}

	// Validate file
	if err := h.validator.ValidateFile(file); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.ErrInvalidInput, err.Error(), nil, utils.GenerateRequestID()))
		return
	}

	// Save file
	filePath, err := h.storage.SaveFile(file, "products", "images")
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(utils.ErrInternalServer, "Failed to save file", nil, utils.GenerateRequestID()))
		return
	}

	// Process image if it's an image file
	fileType := file.Header.Get("Content-Type")
	if strings.HasPrefix(fileType, "image/") {
		if err := h.processor.ProcessImage(filepath.Join(h.storage.BasePath, filePath)); err != nil {
			// Log error but don't fail the upload
			fmt.Printf("Warning: Failed to process image: %v\n", err)
		}
	}

	response := models.UploadResponse{
		Success:      true,
		FileName:     file.Filename,
		OriginalName: file.Filename,
		FileType:     fileType,
		FileSize:     file.Size,
		URL:          fmt.Sprintf("/static/products/images/full/%s", filepath.Base(filePath)),
		Message:      "Product image uploaded successfully",
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(response, "Product image uploaded successfully", utils.GenerateRequestID()))
}

// UploadAvatar handles user avatar uploads
func (h *UploadHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.ErrInvalidInput, "File is required", nil, utils.GenerateRequestID()))
		return
	}

	// Validate file
	if err := h.validator.ValidateFile(file); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.ErrInvalidInput, err.Error(), nil, utils.GenerateRequestID()))
		return
	}

	// Save file
	filePath, err := h.storage.SaveFile(file, "users", "avatars")
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(utils.ErrInternalServer, "Failed to save file", nil, utils.GenerateRequestID()))
		return
	}

	// Process image if it's an image file
	fileType := file.Header.Get("Content-Type")
	if strings.HasPrefix(fileType, "image/") {
		if err := h.processor.ProcessImage(filepath.Join(h.storage.BasePath, filePath)); err != nil {
			// Log error but don't fail the upload
			fmt.Printf("Warning: Failed to process image: %v\n", err)
		}
	}

	response := models.UploadResponse{
		Success:      true,
		FileName:     file.Filename,
		OriginalName: file.Filename,
		FileType:     fileType,
		FileSize:     file.Size,
		URL:          fmt.Sprintf("/static/users/avatars/%s", filepath.Base(filePath)),
		Message:      "Avatar uploaded successfully",
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(response, "Avatar uploaded successfully", utils.GenerateRequestID()))
}

// UploadDocument handles document uploads
func (h *UploadHandler) UploadDocument(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.ErrInvalidInput, "File is required", nil, utils.GenerateRequestID()))
		return
	}

	// Validate file
	if err := h.validator.ValidateFile(file); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.ErrInvalidInput, err.Error(), nil, utils.GenerateRequestID()))
		return
	}

	// Save file
	filePath, err := h.storage.SaveFile(file, "documents", "files")
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(utils.ErrInternalServer, "Failed to save file", nil, utils.GenerateRequestID()))
		return
	}

	response := models.UploadResponse{
		Success:      true,
		FileName:     file.Filename,
		OriginalName: file.Filename,
		FileType:     file.Header.Get("Content-Type"),
		FileSize:     file.Size,
		URL:          fmt.Sprintf("/static/documents/files/%s", filepath.Base(filePath)),
		Message:      "Document uploaded successfully",
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(response, "Document uploaded successfully", utils.GenerateRequestID()))
}

// UploadCategoryBanner handles category banner uploads
func (h *UploadHandler) UploadCategoryBanner(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.ErrInvalidInput, "File is required", nil, utils.GenerateRequestID()))
		return
	}

	// Validate file
	if err := h.validator.ValidateFile(file); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.ErrInvalidInput, err.Error(), nil, utils.GenerateRequestID()))
		return
	}

	// Save file
	filePath, err := h.storage.SaveFile(file, "categories", "banners")
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(utils.ErrInternalServer, "Failed to save file", nil, utils.GenerateRequestID()))
		return
	}

	// Process image if it's an image file
	fileType := file.Header.Get("Content-Type")
	if strings.HasPrefix(fileType, "image/") {
		if err := h.processor.ProcessImage(filepath.Join(h.storage.BasePath, filePath)); err != nil {
			// Log error but don't fail the upload
			fmt.Printf("Warning: Failed to process image: %v\n", err)
		}
	}

	response := models.UploadResponse{
		Success:      true,
		FileName:     file.Filename,
		OriginalName: file.Filename,
		FileType:     fileType,
		FileSize:     file.Size,
		URL:          fmt.Sprintf("/static/categories/banners/%s", filepath.Base(filePath)),
		Message:      "Category banner uploaded successfully",
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(response, "Category banner uploaded successfully", utils.GenerateRequestID()))
}