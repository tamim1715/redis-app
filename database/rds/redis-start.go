package rds

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/khan1507017/redis-app/config"
	"log"
	"sync"
)

var master *redis.Client
var slave [50]*redis.Client
var count int = 0
var mtx sync.Mutex

func InitRedisMaster() error {
	ctx := context.Background()
	master = redis.NewClient(&redis.Options{
		Addr:     config.RedisMasterEndpoint + ":" + config.RedisPort,
		Password: config.RedisPassword, // no password set
		DB:       0,                    // use default DB
	})
	err := master.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		log.Println("Database Connection Error: ", err.Error())
		return err
	} else {
		master.Del(ctx, "key")
	}
	return nil
}
func InitRedisSlave() error {
	for i := 0; i < config.RedisSlaveCount; i++ {
		ctx := context.Background()
		slave[i] = redis.NewClient(&redis.Options{
			Addr:     config.RedisSlaveEndpoints[i] + ":" + config.RedisPort,
			Password: config.RedisPassword, // no password set
			DB:       0,                    // use default DB
		})
		err := slave[i].Set(ctx, "key", "value", 0).Err()
		if err != nil {
			log.Println("Database Connection Error: ", err.Error())
			return err
		} else {
			slave[i].Del(ctx, "key")
		}
	}
	return nil
}

func GetRedisMaster() *redis.Client {
	return master
}
func GetRedisSlave() *redis.Client {
	instance := slave[count]
	mtx.Lock()
	count++
	count = count % config.RedisSlaveCount
	mtx.Unlock()
	return instance
}
