package broker

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/redis/go-redis/v9"
)

// вариант решения для https://github.com/redis/go-redis/issues/2276
func AutoClaim(stream string, group string, consumerName string) ([]redis.XMessage, error) {
	result, _, err := client.XAutoClaim(
		context.Background(),
		&redis.XAutoClaimArgs{
			Stream:   "102",
			Group:    "102",
			Consumer: consumerName,
			MinIdle:  0,
			Start:    "0",
		},
	).Result()
	if err != nil {
		log.Errorf("AutoClaim error:", stream, err)
		return nil, err
	}

	return result, nil
}
