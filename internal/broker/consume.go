package broker

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func Consume() ([]redis.XStream, error) {
	streams, err := Client.XRead(ClientCtx, &redis.XReadArgs{
		Streams: []string{"chat-stream", LastID},
		Block:   0, // ждать новые сообщения
		Count:   10,
	}).Result()

	if err != nil {
		fmt.Println("redis xread error:", err)
		return nil, err
	}

	return streams, nil
}
