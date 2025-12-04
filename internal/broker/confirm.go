package broker

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func ConfirmMesssage(stream string, group string, message redis.XMessage, ctx context.Context) error {
	_, err := client.XAck(ctx, stream, group, message.ID).Result()
	if err != nil {
		log.Print("error during XAck: ", err)
		return err
	}

	log.Printf("ConfirmMesssage: %s by %s \n", stream, group)

	return nil
}
