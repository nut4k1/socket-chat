package broker

import (
	"fmt"

	"github.com/nut4k1/socket-chat/internal/ws"
)

func StartConsumer(hub *ws.Hub) {
	go func() {
		for {
			// читаем новые сообщения из брокера
			streams, err := Consume()
			if err != nil {
				fmt.Println("redis xread error:", err)
				continue
			}

			for _, stream := range streams {
				for _, message := range stream.Messages {
					from, _ := message.Values["from"].(string)
					to, _ := message.Values["to"].(string)
					text, _ := message.Values["message"].(string)

					// формируем JSON для отправки
					msgJSON := fmt.Sprintf(`{"from":"%s","message":"%s"}`, from, text)
					err := hub.SendToUser(to, []byte(msgJSON))
					if err != nil {
						fmt.Println(err)
					}

					UpdateLastID(message.ID) // обновляем lastID
				}
			}
		}
	}()
}
