package server

import (
	"fmt"
	"log"

	"github.com/nut4k1/socket-chat/internal/broker"
	"github.com/nut4k1/socket-chat/internal/config"
	"github.com/nut4k1/socket-chat/internal/http/handlers"
	"github.com/nut4k1/socket-chat/internal/http/middlewares"
	"github.com/nut4k1/socket-chat/internal/ws"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Start(cfg *config.Config) {
	// подключаемся к Redis
	rds := broker.Init(cfg)
	defer rds.Close()

	// инит хаба подключений
	hub := ws.NewHub()

	// go func() {
	// 	for {
	// 		time.Sleep(1 * time.Second)
	// 		fmt.Println("conns:", len(hub.Clients()))
	// 		for id, _ := range hub.Clients() {
	// 			fmt.Println("conn user id:", id)
	// 		}
	// 	}
	// }()

	app := fiber.New()

	wsConfig := websocket.Config{
		ReadBufferSize:  cfg.WS.ReadBufferSize,
		WriteBufferSize: cfg.WS.ReadBufferSize,
	}

	app.Use("/ws", middlewares.WCCheck)
	app.Use("/ws", middlewares.ValidateToken(cfg.Auth.JWTSecret))
	app.Use("/ws", middlewares.DupConn(hub))
	app.Get("/ws", websocket.New(handlers.CreateWCHandler(hub), wsConfig))

	log.Printf("Server started on :%d \n", cfg.Server.Port)
	log.Fatal(app.Listen(fmt.Sprintf(":%d", cfg.Server.Port)))
}
