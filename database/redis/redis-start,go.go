package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/khan1507017/redis-app/config"
)

func InitRedis() error {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.RedisMasterEndpoint + ":" + config.RedisPort,
		Password: config.RedisPassword, // no password set
		DB:       0,                    // use default DB
	})
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		return err
	} else {
		rdb.Del(ctx, "key")
	}
	return nil
}
