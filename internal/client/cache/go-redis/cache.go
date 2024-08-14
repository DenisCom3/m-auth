package goredis

import (
	"context"
	"time"

	"github.com/DenisCom3/m-auth/internal/client/cache"
	"github.com/DenisCom3/m-auth/internal/config"
	redis_cache "github.com/go-redis/cache/v9"
	go_redis "github.com/redis/go-redis/v9"
)

const timeExp = time.Hour

type redisCache struct {
	client *redis_cache.Cache
}

func NewCache(cfg config.Redis) cache.Cache {
	return &redisCache{
		client: redis_cache.New(&redis_cache.Options{
			Redis: go_redis.NewClient(&go_redis.Options{
				PoolSize:              10,
				Addr:                  cfg.Address(),
				Password:              "",
				ConnMaxIdleTime:       cfg.IdleTimeout(),
				ContextTimeoutEnabled: true,
				ConnMaxLifetime:       cfg.ConnectionTimeout(),
				MaxIdleConns:          cfg.MaxIdle(),
			}),
			LocalCache: redis_cache.NewTinyLFU(1000, time.Minute),
		}),
	}
}

func (c *redisCache) Get(ctx context.Context, key string) (any, error) {
	var obj any
	if err := c.client.Get(ctx, key, &obj); err != nil {
		switch {
		case err == go_redis.Nil:
			return nil, cache.ErrNotFound
		default:
			return nil, err
		}
	}
	return obj, nil
}

func (c *redisCache) Set(ctx context.Context, key string, value any) error {
	return c.client.Set(&redis_cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: value,
		TTL:   timeExp,
	})
}
