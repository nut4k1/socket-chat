package main

import (
	"log"
	"os"

	"github.com/nut4k1/socket-chat/internal/broker"
	"github.com/nut4k1/socket-chat/internal/http"
	"github.com/nut4k1/socket-chat/internal/ws"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func main() {
	// подключаемся к Redis
	err := broker.Init()
	if err != nil {
		os.Exit(1)
	}

	// инит хаба подключений
	hub := ws.NewHub()
	// горутина консъюмера
	broker.StartConsumer(hub)

	app := fiber.New()

	// проверка на то что хотят сокет
	// app.Use("/ws", func(c *fiber.Ctx) error {
	// 	if websocket.IsWebSocketUpgrade(c) {
	// 		return c.Next()
	// 	}
	// 	return fiber.ErrUpgradeRequired
	// })

	// проверка на токен
	app.Use("/ws", func(c *fiber.Ctx) error {
		if !websocket.IsWebSocketUpgrade(c) {
			return fiber.ErrUpgradeRequired
		}

		token := c.Query("token")
		if !isValidToken(token) {
			return c.Status(fiber.StatusUnauthorized).SendString("unauthorized")
		}

		return c.Next()
	})

	// Сам WebSocket-обработчик
	app.Get("/ws", websocket.New(http.CreateHandler(hub)))

	log.Println("Server started on :8080")
	log.Fatal(app.Listen(":8080"))
}

func isValidToken(token string) bool {
	return token == "valid-token"
}

// var secret = []byte("secret")
// обработка JWT токена
// func isValidToken(tokenString string) bool {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		return secret, nil
// 	})
// 	if err != nil || !token.Valid {
// 		return false
// 	}
// 	return true
// }
