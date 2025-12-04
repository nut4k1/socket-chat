package handlers

import (
	"encoding/json"
	"log"

	"github.com/nut4k1/socket-chat/internal/broker"

	"github.com/gofiber/websocket/v2"
)

func readAndPublish(userID string, conn *websocket.Conn) error {
	_, msgBytes, err := conn.ReadMessage()
	if err != nil {
		log.Println("conn ReadMessage error:", err)
		return err
	}

	var msg WSMessage
	if err := json.Unmarshal(msgBytes, &msg); err != nil {
		log.Println("invalid message format:", err)
		return nil
	}

	log.Printf("received from %s to %s: %s\n", userID, msg.To, msg.Message)

	err = broker.Publish(msg.To, broker.BrockerMessage{
		From:    userID,
		To:      msg.To,
		Message: msg.Message,
	})
	if err != nil {
		log.Println("broker NewPublish error:", err)
	}

	return nil
}
