package handlers

import (
	"net/http"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/config"
	"github.com/DurgaPratapRajbhar/e-commerce/auth-service/internal/usecase"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/utils"
	"github.com/gin-gonic/gin"
	"log"
	
)

type PasswordHandler struct {
	authUseCase *usecase.AuthUseCase
	useRealImplementation bool
}

func NewPasswordHandler(authUseCase *usecase.AuthUseCase) *PasswordHandler {
	return &PasswordHandler{
		authUseCase: authUseCase,
		useRealImplementation: true,
	}
}

func (h *PasswordHandler) RequestPasswordReset(c *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		validationErrors := utils.ParseValidationErrors(err.Error())
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	if h.useRealImplementation {
		cfg, err := config.Load()
		if err != nil {
			log.Fatalf("⚠️  Failed to load config: %v", err)
		}
		
		data := map[string]interface{}{
			"message": "Password reset email would be sent to " + req.Email,
			"host": cfg.External.SMTPHost,
			"port": cfg.External.SMTPPort,
		}
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(data, "Operation successful", requestID))
	} else {
		data := map[string]string{
			"message": "Password reset email sent successfully",
			"email":   req.Email,
		}
		 
		c.JSON(http.StatusOK, utils.SuccessResponse(data, "Operation successful", utils.GenerateRequestID()))
	}
}

func (h *PasswordHandler) ResetPassword(c *gin.Context) {
	var req struct {
		Token       string `json:"token" binding:"required"`
		NewPassword string `json:"new_password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		validationErrors := utils.ParseValidationErrors(err.Error())
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	if h.useRealImplementation {
		data := map[string]string{
			"message": "Password would be reset for token " + req.Token,
		}
		 
		c.JSON(http.StatusOK, utils.SuccessResponse(data, "Operation successful", utils.GenerateRequestID()))
	} else {
		data := map[string]string{
			"message": "Password reset successfully",
		}
		 
		c.JSON(http.StatusOK, utils.SuccessResponse(data, "Operation successful", utils.GenerateRequestID()))
	}
}