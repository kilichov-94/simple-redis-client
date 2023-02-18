package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis"
)

type Cache struct {
	*redis.Client
}

func New(addr, password string, db int) *Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return &Cache{Client: client}
}

func (c *Cache) Ping(ctx context.Context) error {
	return c.Client.Ping().Err()
}

func (c *Cache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return c.Client.Set(key, value, ttl).Err()
}

func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	return c.Client.Get(key).Result()
}


