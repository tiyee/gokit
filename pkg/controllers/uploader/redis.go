package uploader

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisCache struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisClient(ctx context.Context, client *redis.Client) *RedisCache {
	return &RedisCache{
		client: client,
		ctx:    ctx,
	}
}
func (r *RedisCache) Get(key string) ([]byte, error) {
	return r.client.Get(r.ctx, key).Bytes()
}
func (r *RedisCache) Set(key string, val []byte) error {
	return r.client.Set(r.ctx, key, val, time.Second*3600*24).Err()
}
func (r *RedisCache) Delete(key string) error {
	return r.client.Del(r.ctx, key).Err()
}
