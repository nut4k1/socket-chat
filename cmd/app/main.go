package main

import (
	"os"

	"github.com/nut4k1/socket-chat/internal/config"
	"github.com/nut4k1/socket-chat/internal/http/server"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Overload()

	cfg := config.Load(os.Getenv("CFG_PATH"))

	server.Start(cfg)
}
