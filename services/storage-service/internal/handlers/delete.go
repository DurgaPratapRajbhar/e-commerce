package handlers

import (
	"net/http"

	"github.com/DurgaPratapRajbhar/e-commerce/pkg/utils"
	"github.com/DurgaPratapRajbhar/e-commerce/storage-service/internal/storage"
	"github.com/gin-gonic/gin"
)

type DeleteHandler struct {
	storage *storage.LocalStorage
}

func NewDeleteHandler(storage *storage.LocalStorage) *DeleteHandler {
	return &DeleteHandler{
		storage: storage,
	}
}

func (h *DeleteHandler) DeleteFile(c *gin.Context) {
	// Get category, subcategory, and filename from URL params
	category := c.Param("category")
	if category == "" {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.ErrInvalidInput, "Category is required", nil, utils.GenerateRequestID()))
		return
	}

	subcategory := c.Param("subcategory")
	if subcategory == "" {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.ErrInvalidInput, "Subcategory is required", nil, utils.GenerateRequestID()))
		return
	}

	filename := c.Param("filename")
	if filename == "" {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.ErrInvalidInput, "Filename is required", nil, utils.GenerateRequestID()))
		return
	}

	// Construct the file path
	filePath := filepath.Join(category, subcategory, filename)

	// Check if file exists
	if !h.storage.FileExists(filePath) {
		c.JSON(http.StatusNotFound, utils.ErrorResponse(utils.ErrNotFound, "File not found", nil, utils.GenerateRequestID()))
		return
	}

	// Delete file from storage
	if err := h.storage.DeleteFile(filePath); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(utils.ErrInternalServer, "Failed to delete file", nil, utils.GenerateRequestID()))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(nil, "File deleted successfully", utils.GenerateRequestID()))
}