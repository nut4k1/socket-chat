package broker

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client
var ClientCtx = context.Background()

func Init() error {
	Client = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	_, err := Client.Ping(ClientCtx).Result()
	if err != nil {
		log.Fatal("Redis connection error:", err)
		return err
	}

	return nil
}
