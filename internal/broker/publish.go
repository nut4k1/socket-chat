package broker

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

type BrockerMessage struct {
	From    string
	To      string
	Message string
}

func transformMessage(bm BrockerMessage) map[string]any {
	return map[string]any{
		"from":    bm.From,
		"to":      bm.To,
		"message": bm.Message,
	}
}

func Publish(bm BrockerMessage) error {
	_, err := Client.XAdd(ClientCtx, &redis.XAddArgs{
		Stream: "chat-stream",
		Values: transformMessage(bm),
	}).Result()
	// _, err := Client.XAdd(ClientCtx, &redis.XAddArgs{
	// 	Stream: "chat-stream",
	// 	Values: map[string]interface{}{
	// 		"from":    userID,
	// 		"to":      msg.To,
	// 		"message": msg.Message,
	// 	},
	// }).Result()

	if err != nil {
		fmt.Println("redis xadd error:", err)
		return err
	}
	return nil
}
