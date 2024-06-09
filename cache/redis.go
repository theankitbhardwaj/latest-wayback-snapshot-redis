package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	addr   string
	client redis.Client
}

var ctx = context.Background()

func NewRedisClient(addr string) *RedisClient {

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &RedisClient{addr: addr, client: *rdb}
}

func (r *RedisClient) Setex(key, value string, expiry time.Duration) error {
	return r.client.Set(ctx, key, value, expiry).Err()
}

func (r *RedisClient) Get(key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}
