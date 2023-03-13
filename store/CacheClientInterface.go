package store

import (
	"context"
	"time"
)

type CacheClientInterface interface {
	GetValue(ctx context.Context, key string) (string, error)
	SetValue(ctx context.Context, key string, val interface{}, expiry time.Duration) error
}
