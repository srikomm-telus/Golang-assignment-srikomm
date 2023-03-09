package store

import "time"

type CacheClientInterface interface {
	GetValue(key string) (string, error)
	SetValue(key string, val interface{}, expiry time.Duration) error
}
