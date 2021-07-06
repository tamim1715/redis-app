package rds

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/khan1507017/redis-app/config"
	"log"
)

var rdb *redis.Client

func InitRedisMaster() error {
	ctx := context.Background()
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.RedisMasterEndpoint + ":" + config.RedisPort,
		Password: config.RedisPassword, // no password set
		DB:       0,                    // use default DB
	})
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		log.Println("Database Connection Error: ", err.Error())
		return err
	} else {
		rdb.Del(ctx, "key")
	}
	return nil
}
func GetRedisMaster() *redis.Client {
	return rdb
}
