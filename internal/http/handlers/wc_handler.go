package handlers

import (
	"context"
	"log"

	"github.com/nut4k1/socket-chat/internal/ws"

	"github.com/gofiber/websocket/v2"
)

type WSMessage struct {
	To      string `json:"to"`
	Message string `json:"message"`
}

type HubInterface interface {
	Register(*ws.Client)
	Unregister(string)
	CheckClient(string) bool
	SendToUser(string, []byte) error
}

func CreateWCHandler(hub HubInterface) func(c *websocket.Conn) {
	return func(c *websocket.Conn) {
		userID := c.Query("user_id")

		client := ws.NewClient(userID, c)
		hub.Register(client)
		ctx, cancel := context.WithCancel(context.Background())

		defer func() {
			cancel()
			hub.Unregister(userID)
			c.Close()
			log.Println("Connection closed:", userID)
		}()

		go consumeAndWrite(hub, userID, ctx)

		for {
			err := readAndPublish(userID, c)
			if err != nil {
				log.Println("readAndPublish crushed:", err)
				return
			}
		}
	}
}
