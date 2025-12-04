package ws

import (
	"fmt"
	"log"
	"sync"

	"github.com/gofiber/websocket/v2"
)

type Hub struct {
	clients map[string]*Client
	mu      sync.RWMutex
}

func (h *Hub) Clients() map[string]*Client {
	return h.clients
}

func (h *Hub) Shutdown() error {
	log.Println("Shutdown Hub")
	for userID, client := range h.clients {
		if err := client.Conn.Close(); err != nil {
			log.Println("userID:", userID, "conn close error")
		}
		log.Println("userID:", userID, "conn closed")
	}

	return nil
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[string]*Client),
	}
}

func (h *Hub) CheckClient(userID string) bool {
	h.mu.RLock()
	_, ok := h.clients[userID]
	h.mu.RUnlock()
	return ok
}

func (h *Hub) Register(client *Client) {
	h.mu.Lock()
	h.clients[client.UserID] = client
	h.mu.Unlock()

	log.Println("User Register:", client.UserID)
}

func (h *Hub) Unregister(userID string) {
	h.mu.Lock()
	delete(h.clients, userID)
	h.mu.Unlock()

	log.Println("User Unregister:", userID)
}

func (h *Hub) SendToUser(to string, message []byte) error {
	h.mu.RLock()
	client, ok := h.clients[to]
	h.mu.RUnlock()

	if !ok {
		log.Println("User offline:", to)
		return fmt.Errorf("user offline")
	}

	return client.Conn.WriteMessage(websocket.TextMessage, message)
}
