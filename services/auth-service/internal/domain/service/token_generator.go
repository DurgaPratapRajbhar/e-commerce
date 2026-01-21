package service

import (
	"crypto/rand"
	"encoding/hex"
)

type TokenGenerator interface {
	GenerateToken(length int) (string, error)
}

type RandomTokenGenerator struct{}

func NewRandomTokenGenerator() *RandomTokenGenerator {
	return &RandomTokenGenerator{}
}

func (g *RandomTokenGenerator) GenerateToken(length int) (string, error) {
	bytes := make([]byte, length/2)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}