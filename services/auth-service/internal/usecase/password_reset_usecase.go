package usecase

import (
	"context"
	"time"

	"github.com/DurgaPratapRajbhar/e-commerce/auth-service/internal/domain/entity"
	"github.com/DurgaPratapRajbhar/e-commerce/auth-service/internal/domain/repository"
	"github.com/DurgaPratapRajbhar/e-commerce/auth-service/internal/domain/service"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/utils"
)

type PasswordResetUseCase struct {
	userRepo    repository.UserRepository
	tokenRepo   repository.PasswordResetRepository
	tokenGen    service.TokenGenerator
	emailSender service.EmailService
	tokenExpiry time.Duration
}

func NewPasswordResetUseCase(
	userRepo repository.UserRepository,
	tokenRepo repository.PasswordResetRepository,
	tokenGen service.TokenGenerator,
	emailSender service.EmailService,
	tokenExpiry time.Duration,
) *PasswordResetUseCase {
	return &PasswordResetUseCase{
		userRepo:    userRepo,
		tokenRepo:   tokenRepo,
		tokenGen:    tokenGen,
		emailSender: emailSender,
		tokenExpiry: tokenExpiry,
	}
}

func (uc *PasswordResetUseCase) RequestPasswordReset(ctx context.Context, email string) error {
	user, err := uc.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil
	}

	token, err := uc.tokenGen.GenerateToken(32)
	if err != nil {
		return err
	}

	resetToken := &entity.PasswordResetToken{
		UserID:    user.ID,
		Token:     token,
		ExpiresAt: time.Now().Add(uc.tokenExpiry),
		CreatedAt: time.Now(),
		IsUsed:    false,
	}

	if err := uc.tokenRepo.Create(ctx, resetToken); err != nil {
		return err
	}

	return uc.emailSender.SendPasswordResetEmail(user.Email, token)
}

func (uc *PasswordResetUseCase) ResetPassword(ctx context.Context, token, newPassword string) error {
	resetToken, err := uc.tokenRepo.FindByToken(ctx, token)
	if err != nil {
		return err
	}

	if resetToken == nil {
		return utils.GetError(utils.ErrInvalidToken)
	}

	if time.Now().After(resetToken.ExpiresAt) {
		return utils.GetError(utils.ErrTokenExpired)
	}

	if resetToken.IsUsed {
		return utils.GetError(utils.ErrInvalidToken)
	}

	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}

	user, err := uc.userRepo.FindByID(ctx, resetToken.UserID)
	if err != nil {
		return err
	}

	user.PasswordHash = hashedPassword
	user.UpdatedAt = time.Now()
	if err := uc.userRepo.Update(ctx, user); err != nil {
		return err
	}

	if err := uc.tokenRepo.MarkAsUsed(ctx, resetToken.ID); err != nil {
		return err
	}

	return nil
}