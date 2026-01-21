package token

import (
	"time"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/utils"
	"github.com/DurgaPratapRajbhar/e-commerce/auth-service/internal/domain/service"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/jwt"
)

type jwtService struct {
	secretKey   string
	expiryHours int
}

func NewJWTService(secretKey string, expiryHours int) service.TokenService {
	return &jwtService{
		secretKey:   secretKey,
		expiryHours: expiryHours,
	}
}

func (s *jwtService) GenerateToken(userID int64, email, role string) (string, error) {
	// Use the shared JWT package's GenerateToken function
	// Note: The shared package uses uint for userID, but our interface uses int64
	token, err := jwt.GenerateToken(uint(userID), email, role)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *jwtService) ValidateToken(tokenString string) (*service.TokenClaims, error) {
	// Parse the token with the shared JWT package
	token, err := jwt.ValidateToken(tokenString)
	if err != nil {
		return nil, utils.GetError(utils.ErrInvalidToken)
	}

	// Convert the shared JWT claims to our domain's TokenClaims
	claims := &service.TokenClaims{
		UserID: int64(token.UserID),
		Email:  token.Email,
		Role:   token.Role,
	}
	
	return claims, nil
}

func (s *jwtService) RefreshToken(tokenString string) (string, error) {
	// Validate the existing token first
	claims, err := s.ValidateToken(tokenString)
	if err != nil {
		return "", err
	}
	
	// Generate a new token with the same claims
	return s.GenerateToken(claims.UserID, claims.Email, claims.Role)
}

func (s *jwtService) GetTokenExpiry() time.Duration {
	return time.Duration(s.expiryHours) * time.Hour
}