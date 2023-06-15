package core

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gocli/config"
	"log"
	"time"
)

var Redis = initRedis()

// initRedis 初始化redis客户端
func initRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Config.RedisAddr,
		Password: config.Config.RedisAuth,
		DB:       config.Config.RedisDb,
		PoolSize: config.Config.RedisPoolSize,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal("initRedis client.Ping err: ", err)
	}
	return client
}
