package broker

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func Consume(stream string, group string, consumerName string, ctx context.Context) ([]redis.XStream, error) {
	// “ctx doesn't cancel network read” -_-
	// 	go-redis не может прервать BLOCK-задачу через ctx, потому что:
	// чтение из TCP-сокета нельзя “прервать” обычным контекстом
	// нужен либо timeout, либо ручное закрытие соединения
	// https://github.com/redis/go-redis/issues/2276
	result, err := client.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    group,
		Consumer: consumerName,
		Streams:  []string{stream, ">"},
		Count:    10,
		Block:    5 * time.Second,
	}).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		log.Println("redis xreadgroup error:", err)
		return nil, err
	}
	log.Printf("Consume: XReadGroup from %s by group %s cons %s \n", stream, group, consumerName)

	return result, nil
}
