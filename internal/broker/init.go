package broker

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client
var clientCtx = context.Background()

type redisConfig interface {
	RedisAddr() string
	RedisPassword() string
	RedisDB() int
}

func Init(cfg redisConfig) *redis.Client {
	client = redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr(),
		Password: cfg.RedisPassword(),
		DB:       cfg.RedisDB(),
	})

	_, err := client.Ping(clientCtx).Result()
	if err != nil {
		log.Println("Redis connection error:", err)
	}

	return client
}
