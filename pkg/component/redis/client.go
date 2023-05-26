package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/tiyee/gokit/pkg/consts"
)

var RedisClient *redis.Client

func InitRedis() error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     consts.RedisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	var ctx = context.Background()
	return RedisClient.Ping(ctx).Err()

}
