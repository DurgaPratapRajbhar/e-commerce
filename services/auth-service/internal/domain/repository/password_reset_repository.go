package repository

import (
	"context"
	"github.com/DurgaPratapRajbhar/e-commerce/auth-service/internal/domain/entity"
)

type PasswordResetRepository interface {
	Create(ctx context.Context, token *entity.PasswordResetToken) error
	FindByToken(ctx context.Context, token string) (*entity.PasswordResetToken, error)
	MarkAsUsed(ctx context.Context, id int64) error
	DeleteExpired(ctx context.Context) error
}