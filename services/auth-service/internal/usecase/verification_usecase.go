package usecase

import (
	"context"
	"time"

	"github.com/DurgaPratapRajbhar/e-commerce/auth-service/internal/domain/entity"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/utils"
	"github.com/DurgaPratapRajbhar/e-commerce/auth-service/internal/domain/repository"
	"github.com/DurgaPratapRajbhar/e-commerce/auth-service/internal/domain/service"
)

type VerificationUseCase struct {
	userRepo    repository.UserRepository
	tokenRepo   repository.VerificationRepository
	tokenGen    service.TokenGenerator
	emailSender service.EmailService
	tokenExpiry time.Duration
}

func NewVerificationUseCase(
	userRepo repository.UserRepository,
	tokenRepo repository.VerificationRepository,
	tokenGen service.TokenGenerator,
	emailSender service.EmailService,
	tokenExpiry time.Duration,
) *VerificationUseCase {
	return &VerificationUseCase{
		userRepo:    userRepo,
		tokenRepo:   tokenRepo,
		tokenGen:    tokenGen,
		emailSender: emailSender,
		tokenExpiry: tokenExpiry,
	}
}

func (uc *VerificationUseCase) RequestEmailVerification(ctx context.Context, userID int64) error {
	user, err := uc.userRepo.FindByID(ctx, userID)
	if err != nil {
		return err
	}

	token, err := uc.tokenGen.GenerateToken(32)
	if err != nil {
		return err
	}

	verificationToken := &entity.VerificationToken{
		UserID:    user.ID,
		Token:     token,
		ExpiresAt: time.Now().Add(uc.tokenExpiry),
		CreatedAt: time.Now(),
		IsUsed:    false,
	}

	if err := uc.tokenRepo.Create(ctx, verificationToken); err != nil {
		return err
	}

	return uc.emailSender.SendVerificationEmail(user.Email, token)
}

func (uc *VerificationUseCase) VerifyEmail(ctx context.Context, token string) error {
	verificationToken, err := uc.tokenRepo.FindByToken(ctx, token)
	if err != nil {
		return err
	}

	if verificationToken == nil {
		return utils.GetError(utils.ErrInvalidToken)
	}

	if time.Now().After(verificationToken.ExpiresAt) {
		return utils.GetError(utils.ErrTokenExpired)
	}

	if verificationToken.IsUsed {
		return utils.GetError(utils.ErrInvalidToken)
	}

	user, err := uc.userRepo.FindByID(ctx, verificationToken.UserID)
	if err != nil {
		return err
	}

	user.IsActive = true
	user.UpdatedAt = time.Now()
	if err := uc.userRepo.Update(ctx, user); err != nil {
		return err
	}

	if err := uc.tokenRepo.MarkAsUsed(ctx, verificationToken.ID); err != nil {
		return err
	}

	return nil
}