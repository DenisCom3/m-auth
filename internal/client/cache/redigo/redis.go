package redigo

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/DenisCom3/m-auth/internal/client/cache"
	"github.com/DenisCom3/m-auth/internal/config"
	"github.com/gomodule/redigo/redis"
)

var _ cache.Cache = (*client)(nil)

type handler func(ctx context.Context, conn redis.Conn) error

type client struct {
	pool   *redis.Pool
	config config.Redis
}

func New(pool *redis.Pool, cfg config.Redis) cache.Cache {
	return &client{
		pool:   pool,
		config: cfg,
	}
}

func (c *client) HashSet(ctx context.Context, key string, values interface{}) error {
	err := c.execute(ctx, func(ctx context.Context, conn redis.Conn) error {
		_, err := conn.Do("HSET", redis.Args{key}.AddFlat(values)...)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *client) Set(ctx context.Context, key string, value interface{}) error {
	err := c.execute(ctx, func(ctx context.Context, conn redis.Conn) error {
		_, err := conn.Do("SET", redis.Args{key}.Add(value)...)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *client) HGetAll(ctx context.Context, key string) (map[string]any, error) {
	var values []interface{}
	err := c.execute(ctx, func(ctx context.Context, conn redis.Conn) error {
		var errEx error
		values, errEx = redis.Values(conn.Do("HGETALL", key))
		if errEx != nil {
			return errEx
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	result := make(map[string]any)
	if condition := len(values)%2 != 0; condition {
		return nil, errors.New("redigo: HGET. invalid number of values")
	}
	for i := 0; i < len(values); i += 2 {
		result[values[i].(string)] = values[i+1]
	}

	return result, nil
}

func (c *client) Get(ctx context.Context, key string) (interface{}, error) {
	var val any
	err := c.execute(ctx, func(ctx context.Context, conn redis.Conn) error {
		var ErrEx error
		val, ErrEx = conn.Do("GET", key)

		if ErrEx != nil {
			return ErrEx
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return val, err

}

func (c *client) Expire(ctx context.Context, key string, expiration time.Duration) error {
	err := c.execute(ctx, func(ctx context.Context, conn redis.Conn) error {
		_, err := conn.Do("EXPIRE", key, int(expiration.Seconds()))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *client) Ping(ctx context.Context) error {

	err := c.execute(ctx, func(ctx context.Context, conn redis.Conn) error {
		_, err := conn.Do("PING")
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (c *client) execute(ctx context.Context, handler handler) error {
	conn, err := c.getConnect(ctx)
	if err != nil {
		return err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Printf("failed to close redigo connection: %v\n", err)
		}
	}()

	err = handler(ctx, conn)
	if err != nil {
		return err
	}

	return nil
}

func (c *client) getConnect(ctx context.Context) (redis.Conn, error) {
	getConnTimeoutCtx, cancel := context.WithTimeout(ctx, c.config.ConnectionTimeout())
	defer cancel()

	conn, err := c.pool.GetContext(getConnTimeoutCtx)
	if err != nil {
		log.Printf("failed to get redigo connection: %v\n", err)

		_ = conn.Close()
		return nil, err
	}

	return conn, nil
}
