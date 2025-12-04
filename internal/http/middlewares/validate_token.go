package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte{}

func keyFunc(token *jwt.Token) (any, error) {
	if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	return secret, nil
}

func isValidToken(tokenString string, userID string) bool {
	token, err := jwt.Parse(tokenString, keyFunc)
	if err != nil || !token.Valid {
		return false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false
	}

	tokenUserID, ok := claims["user_id"].(string)
	return ok && tokenUserID == userID
}

func ValidateToken(s string) func(ctx *fiber.Ctx) error {
	secret = []byte(s)

	return func(ctx *fiber.Ctx) error {
		token := ctx.Query("token")
		userID := ctx.Query("user_id")
		if !isValidToken(token, userID) {
			return ctx.Status(fiber.StatusUnauthorized).SendString("Unauthorized: invalid token")
		}

		return ctx.Next()
	}
}

// func ValidateToken(ctx *fiber.Ctx) error {
// 	token := ctx.Query("token")
// 	userID := ctx.Query("user_id")
// 	if !isValidToken(token, userID) {
// 		return ctx.Status(fiber.StatusUnauthorized).SendString("Unauthorized: invalid token")
// 	}

// 	return ctx.Next()
// }
