package handlers

import (
	"net/http"
	"strings"
	"github.com/DurgaPratapRajbhar/e-commerce/auth-service/internal/usecase"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/utils"
	"github.com/gin-gonic/gin"
)

type TokenHandler struct {
	tokenUseCase *usecase.TokenUseCase
}

func NewTokenHandler(tokenUseCase *usecase.TokenUseCase) *TokenHandler {
	return &TokenHandler{tokenUseCase: tokenUseCase}
}

func (h *TokenHandler) ValidateToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse(
			utils.ErrUnauthorized,
			"Unauthorized",
			"missing authorization header",
			requestID,
		))
		return
	}
	
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	
	claims, err := h.tokenUseCase.ValidateToken(c.Request.Context(), tokenString)
	if err != nil {
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse(
			utils.ErrInvalidToken,
			"Unauthorized",
			err.Error(),
			requestID,
		))
		return
	}
	
	data := map[string]interface{}{
		"valid":  true,
		"claims": claims,
	}
	requestID := utils.GenerateRequestID()
	c.JSON(http.StatusOK, utils.SuccessResponse(data, "Token validation successful", requestID))
}

func (h *TokenHandler) RefreshToken(c *gin.Context) {
	var req struct {
		Token string `json:"token" binding:"required"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		validationErrors := utils.ParseValidationErrors(err.Error())
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}
	
	newToken, err := h.tokenUseCase.RefreshToken(c.Request.Context(), req.Token)
	if err != nil {
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse(
			utils.ErrInvalidToken,
			"Unauthorized",
			err.Error(),
			requestID,
		))
		return
	}
	
	data := map[string]string{
		"token": newToken,
	}
	requestID := utils.GenerateRequestID()
	c.JSON(http.StatusOK, utils.SuccessResponse(data, "Token validation successful", requestID))
}