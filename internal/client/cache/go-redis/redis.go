package goredis

import (
	"context"
	"time"

	"github.com/DenisCom3/m-auth/internal/client/cache"
	"github.com/DenisCom3/m-auth/internal/config"
	redis "github.com/redis/go-redis/v9"
)

var _ cache.Cache = (*client)(nil)

type client struct {
	client *redis.Client
}

func NewClient(cfg config.Redis) cache.Cache {
	return &client{
		client: redis.NewClient(&redis.Options{
			PoolSize:              10,
			Addr:                  cfg.Address(),
			Password:              "",
			ConnMaxIdleTime:       cfg.IdleTimeout(),
			ContextTimeoutEnabled: true,
			ConnMaxLifetime:       cfg.ConnectionTimeout(),
			MaxIdleConns:          cfg.MaxIdle(),
		}),
	}
}

func (c *client) Ping(ctx context.Context) error {
	return c.client.Ping(ctx).Err()
}

func (c *client) Set(ctx context.Context, key string, value interface{}) error {
	return c.client.Set(ctx, key, value, 0).Err()
}

func (c *client) Get(ctx context.Context, key string) (any, error) {
	var obj any
	err := c.client.Get(ctx, key).Scan(&obj)

	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *client) HGetAll(ctx context.Context, key string) (map[string]interface{}, error) {

	values, err := c.client.HGetAll(ctx, key).Result()

	if err != nil {
		return nil, err
	}

	result := make(map[string]any, len(values))
	for k, v := range values {
		result[k] = v
	}
	return result, nil
}

func (c *client) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return c.client.Expire(ctx, key, expiration).Err()
}

func (c *client) HashSet(ctx context.Context, key string, values interface{}) error {
	return c.client.HSet(ctx, key, values).Err()
}
