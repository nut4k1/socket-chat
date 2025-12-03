package http

import (
	"encoding/json"
	"fmt"

	"github.com/nut4k1/socket-chat/internal/broker"
	"github.com/nut4k1/socket-chat/internal/ws"

	"github.com/gofiber/websocket/v2"
)

type WSMessage struct {
	To      string `json:"to"`
	Message string `json:"message"`
}

func CreateHandler(hub *ws.Hub) func(c *websocket.Conn) {
	return func(c *websocket.Conn) {
		// 1. Авторизация по токену и сохранение коннекта по user_id
		// token := c.Query("token")
		// if token != "valid-token" {
		// 	c.WriteMessage(websocket.TextMessage, []byte("unauthorized"))
		// 	c.Close()
		// 	return
		// }
		userID := c.Query("user_id")

		// 2. Сохраняем соединение
		client := ws.NewClient(userID, c)
		hub.Register(client)
		defer func() {
			hub.Unregister(userID)
			c.Close()
		}()

		// 3. Чтение сообщений от клиента
		for {
			_, msgBytes, err := c.ReadMessage()
			if err != nil {
				fmt.Println("read error:", err)
				return
			}

			// пересобираем в наш стракт
			var msg WSMessage
			if err := json.Unmarshal(msgBytes, &msg); err != nil {
				fmt.Println("invalid message format:", err)
				continue
			}

			fmt.Printf("received from %s to %s: %s\n", userID, msg.To, msg.Message)

			// Кладём сообщение в Redis Stream
			err = broker.Publish(broker.BrockerMessage{
				From:    userID,
				To:      msg.To,
				Message: msg.Message,
			})
			if err != nil {
				fmt.Println("redis xadd error:", err)
			}
		}
	}
}
