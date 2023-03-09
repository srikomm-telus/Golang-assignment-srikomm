package store

import (
	"context"
	"time"
)

type CacheClientInterface interface {
	GetValue(key string, ctx context.Context) (string, error)
	SetValue(key string, ctx context.Context, val interface{}, expiry time.Duration) error
}
