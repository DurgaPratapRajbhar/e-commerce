package handlers

import (
	"net/http"
	"strconv"
	"github.com/DurgaPratapRajbhar/e-commerce/auth-service/internal/usecase"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/utils"
	"github.com/gin-gonic/gin"
	 
	
)

type AuthHandler struct {
	authUseCase *usecase.AuthUseCase
}

func NewAuthHandler(authUseCase *usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{authUseCase: authUseCase}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req usecase.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(
			utils.ErrValidationFailed,
			"Invalid request data",
			err.Error(),
			requestID,
		))
		return
	}
	
	user, err := h.authUseCase.Register(c.Request.Context(), req)
	if err != nil {
		requestID := utils.GenerateRequestID()
		
		if err.Error() == "user already exists" {
			c.JSON(http.StatusConflict, utils.ErrorResponse(
				utils.ErrAlreadyExists,
				"User already exists",
				err.Error(),
				requestID,
			))
		} else {
			c.JSON(http.StatusBadRequest, utils.ErrorResponse(
				utils.ErrValidationFailed,
				"Registration failed",
				err.Error(),
				requestID,
			))
		}
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    user,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req usecase.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(
			utils.ErrValidationFailed,
			"Invalid request data",
			err.Error(),
			requestID,
		))
		return
	}
	
	response, err := h.authUseCase.Login(c.Request.Context(), req)
	if err != nil {
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse(
			utils.ErrInvalidCredentials,
			"Login failed",
			err.Error(),
			requestID,
		))
		return
	}
	
	
	token := response.Token
	// Cookie management is handled by the API Gateway
	// The auth service only returns the token for the gateway to handle
	
	
	responseData := map[string]interface{}{
		"success": true,
		"message": "Login successful",
		"user":    response.User,
		"token":   token, // Include token for API Gateway to handle cookie management
	}
	
	c.JSON(http.StatusOK, responseData)
}

func (h *AuthHandler) GetMe(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse(
			utils.ErrUnauthorized,
			"Unauthorized",
			nil,
			requestID,
		))
		return
	}
	
	// Handle both uint (from JWT claims) and int64 (from headers) types
	var userIDInt64 int64
	switch v := userID.(type) {
	case uint:
		userIDInt64 = int64(v)
	case int64:
		userIDInt64 = v
	case float64: // From JWT payload parsing
		userIDInt64 = int64(v)
	default:
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(
			utils.ErrInternalServer,
			"Invalid user ID type",
			nil,
			requestID,
		))
		return
	}
	
	user, err := h.authUseCase.GetUserByID(c.Request.Context(), userIDInt64)
	if err != nil {
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusNotFound, utils.ErrorResponse(
			utils.ErrNotFound,
			"User not found",
			err.Error(),
			requestID,
		))
		return
	}
	
	requestID := utils.GenerateRequestID()
	c.JSON(http.StatusOK, utils.SuccessResponse(user, "User data retrieved", requestID))
}

func (h *AuthHandler) Logout(c *gin.Context) {
	
	c.SetCookie(
		"auth_token",          // cookie name
		"",                    // empty value to clear
		-3600,                 // negative max age to delete immediately
		"/",                  // path
		"",                   // domain
		false,                 // secure (set to true in production with HTTPS)
		true,                  // httpOnly
	)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Logged out successfully",
	})
}

func (h *AuthHandler) GetAdminUsers(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")
	
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		pageInt = 1
	}
	
	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt < 1 || limitInt > 100 {
		limitInt = 10
	}
	
	offset := (pageInt - 1) * limitInt
	
	users, err := h.authUseCase.GetAllUsers(c.Request.Context(), offset, limitInt)
	if err != nil {
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(
			utils.ErrInternalServer,
			"Failed to fetch users",
			err.Error(),
			requestID,
		))
		return
	}
	
	
	totalCount, err := h.authUseCase.GetTotalUserCount(c.Request.Context())
	if err != nil {
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(
			utils.ErrInternalServer,
			"Failed to get total user count",
			err.Error(),
			requestID,
		))
		return
	}
	
	pagination := utils.NewPagination(pageInt, limitInt, totalCount)
	paginationResponse := utils.PaginationResponse{
		Data:       users,
		Pagination: pagination,
	}
	
	requestID := utils.GenerateRequestID()
	c.JSON(http.StatusOK, utils.SuccessResponse(paginationResponse, "Users retrieved successfully", requestID))
}