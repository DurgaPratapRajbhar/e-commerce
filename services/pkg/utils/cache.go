package utils

import (
	"encoding/json"
	"fmt"
	"time"
)

// CacheKey generates cache key with prefix
func CacheKey(prefix string, id interface{}) string {
	return fmt.Sprintf("%s:%v", prefix, id)
}

// CacheKeyWithParams generates cache key with multiple params
func CacheKeyWithParams(prefix string, params ...interface{}) string {
	key := prefix
	for _, param := range params {
		key += fmt.Sprintf(":%v", param)
	}
	return key
}

// SerializeForCache converts struct to JSON string for caching
func SerializeForCache(data interface{}) (string, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// DeserializeFromCache converts JSON string back to struct
func DeserializeFromCache(data string, target interface{}) error {
	return json.Unmarshal([]byte(data), target)
}

// GetCacheTTL returns cache TTL based on data type
func GetCacheTTL(dataType string) time.Duration {
	ttlMap := map[string]time.Duration{
		"product":  30 * time.Minute,
		"category": 1 * time.Hour,
		"user":     15 * time.Minute,
		"cart":     24 * time.Hour,
		"config":   24 * time.Hour,
	}

	if ttl, exists := ttlMap[dataType]; exists {
		return ttl
	}
	return 5 * time.Minute // default
}