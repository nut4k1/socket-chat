package broker

import (
	"log"

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

func Publish(stream string, bm BrockerMessage) error {
	_, err := client.XAdd(clientCtx, &redis.XAddArgs{
		Stream: stream, // должен создать стрим если его нет
		Values: transformMessage(bm),
	}).Result()

	log.Printf("Publish: add msg to %s \n", stream)

	if err != nil {
		log.Println("redis xadd error:", err)
		return err
	}
	return nil
}
