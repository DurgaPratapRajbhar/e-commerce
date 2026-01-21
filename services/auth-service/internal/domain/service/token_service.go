package service

import "time"

type TokenClaims struct {
	UserID   int64  `json:"user_id"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	IssuedAt int64  `json:"iat"`
	ExpiresAt int64 `json:"exp"`
}

type TokenService interface {
	GenerateToken(userID int64, email, role string) (string, error)
	ValidateToken(token string) (*TokenClaims, error)
	RefreshToken(token string) (string, error)
	GetTokenExpiry() time.Duration
}