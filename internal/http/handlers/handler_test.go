package handlers

import (
	"math/rand"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/fasthttp/websocket"
	"github.com/gofiber/fiber/v2"

	w2 "github.com/gofiber/websocket/v2"
	"github.com/nut4k1/socket-chat/internal/broker"
	"github.com/nut4k1/socket-chat/internal/ws"
)

type MockHub struct {
	RegisterCalled   bool
	UnregisterCalled bool
}

func (m *MockHub) Register(_ *ws.Client)               { m.RegisterCalled = true }
func (m *MockHub) Unregister(s string)                 { m.UnregisterCalled = true }
func (m *MockHub) CheckClient(s string) bool           { return true }
func (m *MockHub) SendToUser(s string, b []byte) error { return nil }

func generate_body_key() string {
	key := make([]byte, 8)
	for i := range 8 {
		key[i] = byte(rand.Intn(0x07e-0x21+1) + 0x21) // случайный символ от 0x21 до 0x7e
	}

	return string(key)
}

func TestWSHandler(t *testing.T) {
	s := miniredis.RunT(t)
	client := broker.Init(&broker.FakeRedisConfig{Addr: s.Addr()})
	defer client.Close()

	mockHub := &MockHub{}
	app := fiber.New()
	app.Get("/ws", w2.New(CreateWCHandler(mockHub)))

	go func() {
		if err := app.Listen(":8080"); err != nil {
		}
	}()

	time.Sleep(100 * time.Millisecond)

	url := "ws://localhost:8080/ws?user_id=test"
	conn, _, _ := websocket.DefaultDialer.Dial(url, nil)
	defer conn.Close()

	req := httptest.NewRequest("GET", "/ws", nil)
	req.Header.Set("Connection", "Upgrade")
	req.Header.Set("Upgrade", "websocket")
	req.Header.Set("Sec-Websocket-Key", generate_body_key())
	req.Header.Set("Sec-Websocket-Version", "13")
	_, _ = app.Test(req, -1)
	if !mockHub.RegisterCalled {
		t.Fatal("hub.RegisterCalled was not called")
	}
}
