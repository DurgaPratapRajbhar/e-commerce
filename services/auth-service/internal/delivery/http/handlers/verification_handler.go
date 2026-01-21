package handlers

import (
	"net/http"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/config"
	"github.com/DurgaPratapRajbhar/e-commerce/auth-service/internal/usecase"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/utils"
	"github.com/gin-gonic/gin"
	"log"
	
)

type VerificationHandler struct {
	authUseCase *usecase.AuthUseCase
	useRealImplementation bool
}

func NewVerificationHandler(authUseCase *usecase.AuthUseCase) *VerificationHandler {
	return &VerificationHandler{
		authUseCase: authUseCase,
		useRealImplementation: true,
	}
}

func (h *VerificationHandler) RequestEmailVerification(c *gin.Context) {
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
			"message": "Verification email would be sent to " + req.Email,
			"host": cfg.External.SMTPHost,
			"port": cfg.External.SMTPPort,
		}
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(data, "Operation successful", requestID))
	} else {
		data := map[string]string{
			"message": "Verification email sent successfully",
			"email":   req.Email,
		}
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(data, "Operation successful", requestID))
	}
}

func (h *VerificationHandler) VerifyEmail(c *gin.Context) {
	var req struct {
		Token string `json:"token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		validationErrors := utils.ParseValidationErrors(err.Error())
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	if h.useRealImplementation {
		data := map[string]string{
			"message": "Email would be verified for token " + req.Token,
		}
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(data, "Operation successful", requestID))
	} else {
		data := map[string]string{
			"message": "Email verified successfully",
		}
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(data, "Operation successful", requestID))
	}
}