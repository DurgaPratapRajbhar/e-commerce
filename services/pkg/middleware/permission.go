package middleware

import (
	"net/http"

	"github.com/DurgaPratapRajbhar/e-commerce/pkg/utils"
	"github.com/gin-gonic/gin"
)

// PermissionMiddleware checks if the authenticated user has the required permission
func PermissionMiddleware(requiredPerm utils.Permission) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user role from context (set by AuthMiddleware)
		userRole, exists := c.Get("user_role")
		if !exists {
			c.JSON(http.StatusForbidden, utils.ErrorResponse(utils.ErrForbidden, "User role not found in context", nil, utils.GenerateRequestID()))
			c.Abort()
			return
		}

		// Convert role to the expected type
		role, ok := userRole.(string)
		if !ok {
			c.JSON(http.StatusForbidden, utils.ErrorResponse(utils.ErrForbidden, "Invalid user role type", nil, utils.GenerateRequestID()))
			c.Abort()
			return
		}

		// Check if user has the required permission
		if !utils.HasPermission(utils.Role(role), requiredPerm) {
			c.JSON(http.StatusForbidden, utils.ErrorResponse(utils.ErrForbidden, "Insufficient permissions", nil, utils.GenerateRequestID()))
			c.Abort()
			return
		}

		// User has required permission, continue
		c.Next()
	}
}