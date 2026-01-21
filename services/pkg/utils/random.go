package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"math/big"
)

// GenerateRandomString generates random string of specified length
func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		b[i] = charset[num.Int64()]
	}
	return string(b)
}

// GenerateRandomNumber generates random number between min and max
func GenerateRandomNumber(min, max int) int {
	diff := max - min
	num, _ := rand.Int(rand.Reader, big.NewInt(int64(diff)))
	return int(num.Int64()) + min
}

// GenerateSecureToken generates cryptographically secure token
func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

// GenerateOTP generates 6-digit OTP
func GenerateOTP() string {
	num := GenerateRandomNumber(100000, 999999)
	return fmt.Sprintf("%06d", num)
}

// GenerateVerificationCode generates verification code
func GenerateVerificationCode(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		b[i] = charset[num.Int64()]
	}
	return string(b)
}