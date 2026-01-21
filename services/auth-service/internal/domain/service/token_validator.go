package service

import (
	"time"
)

type TokenValidator interface {
	IsTokenValid(token string, expiry time.Time) bool
	IsTokenExpired(expiry time.Time) bool
}