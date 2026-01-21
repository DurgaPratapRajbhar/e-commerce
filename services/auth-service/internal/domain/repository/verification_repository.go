package repository

import (
	"context"
	"github.com/DurgaPratapRajbhar/e-commerce/auth-service/internal/domain/entity"
)

type VerificationRepository interface {
	Create(ctx context.Context, token *entity.VerificationToken) error
	FindByToken(ctx context.Context, token string) (*entity.VerificationToken, error)
	MarkAsUsed(ctx context.Context, id int64) error
	DeleteExpired(ctx context.Context) error
}