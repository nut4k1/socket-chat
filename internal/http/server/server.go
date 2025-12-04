package server

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nut4k1/socket-chat/internal/broker"
	"github.com/nut4k1/socket-chat/internal/config"
	"github.com/nut4k1/socket-chat/internal/http/handlers"
	"github.com/nut4k1/socket-chat/internal/http/middlewares"
	"github.com/nut4k1/socket-chat/internal/ws"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Start(cfg *config.Config) {
	rds := broker.Init(cfg)
	defer rds.Close()

	hub := ws.NewHub()
	app := fiber.New()

	wsConfig := websocket.Config{
		ReadBufferSize:  cfg.WS.ReadBufferSize,
		WriteBufferSize: cfg.WS.ReadBufferSize,
	}

	app.Use("/ws", middlewares.WCCheck)
	app.Use("/ws", middlewares.ValidateToken(cfg.Auth.JWTSecret))
	app.Use("/ws", middlewares.DupConn(hub))
	app.Get("/ws", websocket.New(handlers.CreateWCHandler(hub), wsConfig))

	go func() {
		log.Printf("Server started on :%d \n", cfg.Server.Port)
		if err := app.Listen(":8080"); err != nil {
			log.Println("Fiber crushed:", err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := hub.Shutdown(); err != nil {
		log.Println("Hub shutdown error:", err)
	}
	log.Println("app ShutdownWithContext")
	if err := app.ShutdownWithContext(shutdownCtx); err != nil {
		log.Println("Fiber shutdown error:", err)
	}
}
