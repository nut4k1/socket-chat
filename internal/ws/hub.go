package ws

import (
	"fmt"
	"sync"

	"github.com/fasthttp/websocket"
)

type Hub struct {
	clients map[string]*Client
	mu      sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[string]*Client),
	}
}

// Добавить клиента
func (h *Hub) Register(client *Client) {
	h.mu.Lock()
	h.clients[client.UserID] = client
	h.mu.Unlock()

	fmt.Println("User connected:", client.UserID)
}

// Удалить клиента
func (h *Hub) Unregister(userID string) {
	h.mu.Lock()
	delete(h.clients, userID)
	h.mu.Unlock()

	fmt.Println("User disconnected:", userID)
}

// Отправка сообщения конкретному пользователю
func (h *Hub) SendToUser(to string, message []byte) error {
	h.mu.RLock()
	client, ok := h.clients[to]
	h.mu.RUnlock()

	if !ok {
		fmt.Println("User offline:", to)
		return fmt.Errorf("user offline")
	}

	return client.Conn.WriteMessage(websocket.TextMessage, message)
}
