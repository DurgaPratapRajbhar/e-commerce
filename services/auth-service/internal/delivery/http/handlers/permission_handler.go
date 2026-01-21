package handlers

import (
	"net/http"
	"github.com/DurgaPratapRajbhar/e-commerce/auth-service/internal/domain/repository"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/utils"
	"github.com/gin-gonic/gin"
)

type PermissionHandler struct {
	userRepo repository.UserRepository
}

func NewPermissionHandler(userRepo repository.UserRepository) *PermissionHandler {
	return &PermissionHandler{userRepo: userRepo}
}

func (h *PermissionHandler) CheckUserRole(c *gin.Context) {
	userID := c.Param("userID")
	
	data := map[string]string{
		"user_id": userID,
		"message": "In a complete implementation, this would return the user's role from the database",
	}
	requestID := utils.GenerateRequestID()
	c.JSON(http.StatusOK, utils.SuccessResponse(data, "Operation successful", requestID))
}

func (h *PermissionHandler) GetUserPermissions(c *gin.Context) {
	role := c.Query("role")
	if role == "" {
		validationErrors := []utils.ValidationError{
			{Field: "role", Message: "role query parameter is required"},
		}
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	perms := utils.GetRolePermissions(utils.Role(role))
	
	data := map[string]interface{}{
		"role":        role,
		"permissions": perms,
	}
	requestID := utils.GenerateRequestID()
	c.JSON(http.StatusOK, utils.SuccessResponse(data, "Operation successful", requestID))
}

func (h *PermissionHandler) CheckPermission(c *gin.Context) {
	role := c.Query("role")
	permission := c.Query("permission")
	
	if role == "" || permission == "" {
		validationErrors := []utils.ValidationError{}
		if role == "" {
			validationErrors = append(validationErrors, utils.ValidationError{
				Field: "role",
				Message: "role query parameter is required",
			})
		}
		if permission == "" {
			validationErrors = append(validationErrors, utils.ValidationError{
				Field: "permission", 
				Message: "permission query parameter is required",
			})
		}
		
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	hasPerm := utils.HasPermission(utils.Role(role), utils.Permission(permission))
	
	data := map[string]interface{}{
		"role":       role,
		"permission": permission,
		"has_access": hasPerm,
	}
	requestID := utils.GenerateRequestID()
	c.JSON(http.StatusOK, utils.SuccessResponse(data, "Operation successful", requestID))
}