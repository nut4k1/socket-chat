package handlers

import (
	"context"
	"fmt"
	"log"

	"github.com/nut4k1/socket-chat/internal/broker"

	"github.com/redis/go-redis/v9"
)

func processMessage(hub HubInterface, userID string, msg redis.XMessage, ctx context.Context) error {
	from, _ := msg.Values["from"].(string)
	to, _ := msg.Values["to"].(string)
	text, _ := msg.Values["message"].(string)

	log.Printf("Consumer <%s> read msg id: %s\n", userID, msg.ID)

	msgJSON := fmt.Sprintf(`{"from":"%s","message":"%s"}`, from, text)
	err := hub.SendToUser(to, []byte(msgJSON))
	if err != nil {
		log.Println("hub SendToUser error:", err)
		return err
	}

	err = broker.ConfirmMesssage(userID, userID, msg, ctx)
	if err != nil {
		log.Println("error during broker ConfirmMesssage:", err)
		return err
	}

	return nil
}
