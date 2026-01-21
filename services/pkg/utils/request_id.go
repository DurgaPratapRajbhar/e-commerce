package utils

import (
	"crypto/rand"
	"fmt"
	mathRand "math/rand"
	"time"
)

// GenerateRequestID generates a unique request ID
func GenerateRequestID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		// fallback to timestamp if random generation fails
		r := mathRand.New(mathRand.NewSource(time.Now().UnixNano()))
		return fmt.Sprintf("req_%d", r.Uint32())
	}
	
	return fmt.Sprintf("%x-%x-%x-%x-%x", 
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

// GenerateShortRequestID generates a short unique request ID
func GenerateShortRequestID() string {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		// Seed if needed for fallback
		r := mathRand.New(mathRand.NewSource(time.Now().UnixNano()))
		return fmt.Sprintf("req_%d", r.Uint32())
	}
	
	return fmt.Sprintf("%x", b[:8])
}