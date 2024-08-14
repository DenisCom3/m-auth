package cache

import (
	"context"
	"errors"
	"time"
)

type RedisClient interface {
	HashSet(ctx context.Context, key string, values any) error
	Set(ctx context.Context, key string, value any) error
	HGetAll(ctx context.Context, key string) (map[string]any, error)
	Get(ctx context.Context, key string) (any, error)
	Expire(ctx context.Context, key string, expiration time.Duration) error
	Ping(ctx context.Context) error
}

type Cache interface {
	Get(ctx context.Context, key string) (any, error)
	Set(ctx context.Context, key string, value any) error
}

var (
	ErrNotFound      = errors.New("not found from cache")
	ErrTypeAssertion = errors.New("type assertion. error model from cache")
)
