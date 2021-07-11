package rds

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/khan1507017/redis-app/config"
	"log"
	"strconv"
	"sync"
	"time"
)

var master *redis.Client
var slave [50]*redis.Client
var count = 0
var mtx sync.Mutex

func InitRedisMaster() error {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(5))
	master = redis.NewClient(&redis.Options{
		Addr:     config.RedisMasterEndpoint + ":" + config.RedisPort,
		Password: config.RedisPassword,
		DB:       0,
	})
	err := master.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		log.Println("Database Connection Error: ", err.Error())
		return err
	} else {
		master.Del(ctx, "key")
	}
	fmt.Println("master instance loaded")
	return nil
}
func InitRedisSlave() error {
	for i := 0; i < config.RedisSlaveCount; i++ {
		ctx, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(5))
		slave[i] = redis.NewClient(&redis.Options{
			Addr:     config.RedisSlaveEndpoints[i] + ":" + config.RedisPort,
			Password: config.RedisPassword,
			DB:       0,
		})
		if err := slave[i].Ping(ctx).Err(); err != nil {
			return errors.New("error connecting slave: " + strconv.Itoa(i))
		}
	}
	fmt.Println("slave instance loaded: " + strconv.Itoa(config.RedisSlaveCount))
	return nil
}

func GetRedisMaster() *redis.Client {
	return master
}
func GetRedisSlave() *redis.Client {
	if config.RedisSlaveCount == 0 {
		return master
	}
	instance := slave[count]
	mtx.Lock()
	count++
	count = count % config.RedisSlaveCount
	mtx.Unlock()
	return instance
}
