package handlers

import (
	"context"
	"log"
	"time"

	"github.com/nut4k1/socket-chat/internal/broker"
	"github.com/nut4k1/socket-chat/internal/ws"
)

func consumeAndWrite(hub *ws.Hub, userID string, ctx context.Context) {
	err := broker.EnsureGroup(userID, userID, ctx)
	if err != nil {
		log.Println("broker NewEnsureGroup error:", err)
		return
	}

	// сначала проверив PEL и обработаем их
	consumeAndWritePEL(hub, userID, ctx)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			streams, err := broker.Consume(userID, userID, userID, ctx)
			if err != nil {
				log.Println("broker Consume error:", err)
				time.Sleep(1 * time.Second)
				continue
			}

			for _, s := range streams {
				for _, msg := range s.Messages {
					err := processMessage(hub, userID, msg, ctx)
					if err != nil {
						log.Fatal("consumeAndWrite processMessage error:", err)
					}
				}
			}
		}
	}
}
