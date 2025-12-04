package handlers

import (
	"context"
	"log"

	"github.com/nut4k1/socket-chat/internal/broker"
	"github.com/nut4k1/socket-chat/internal/ws"
)

func consumeAndWritePEL(hub *ws.Hub, userID string, ctx context.Context) {
	select {
	case <-ctx.Done():
		return
	default:
		log.Printf("readPEL user id: %s\n", userID)
		msgs, err := broker.AutoClaim(userID, userID, userID)
		if err != nil {
			log.Fatal("broker AutoClaim error:", err)
			return
		}

		for _, msg := range msgs {
			err := processMessage(hub, userID, msg, ctx)
			if err != nil {
				log.Fatal("readPEL processMessage error:", err)
			}
		}
	}
}
