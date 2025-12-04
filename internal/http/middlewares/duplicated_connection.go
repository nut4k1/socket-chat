package middlewares

import (
	"log"

	"github.com/nut4k1/socket-chat/internal/ws"

	"github.com/gofiber/fiber/v2"
)

func DupConn(hub *ws.Hub) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		userID := ctx.Query("user_id")

		if hub.CheckClient(userID) {
			log.Println("DupConn detected")
			return ctx.Status(fiber.StatusUnauthorized).SendString("Unauthorized: user already connected")
		}

		return ctx.Next()
	}
}
