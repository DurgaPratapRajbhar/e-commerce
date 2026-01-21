package jwt

import (
	"errors"
	"fmt"
	// "os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/config"
)

var secretKey = []byte(getSecretKey())

func getSecretKey() string {
	cfg, err := config.Load()
	if err != nil {
		fmt.Println("⚠️  Failed to load config")

	}

	key := cfg.JWT.SecretKey
	if key == "" {
		fmt.Println("⚠️  JWT_SECRET not found, using default")
		return "d9F3kP2xR7mQ8WZL4N6aYtHcB5S0EJUV"
	}
	return key
}

type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateToken creates a new JWT token
func GenerateToken(userID uint, email, role string) (string, error) {
	claims := Claims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// ValidateToken verifies and parses JWT token
func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// RefreshToken generates a new token from existing valid token
func RefreshToken(tokenString string) (string, error) {
	claims, err := ValidateToken(tokenString)
	if err != nil {
		return "", err
	}

	// Generate new token with same claims
	return GenerateToken(claims.UserID, claims.Email, claims.Role)
}
 