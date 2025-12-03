package ws

import "github.com/gofiber/websocket/v2"

type Client struct {
	UserID string
	Conn   *websocket.Conn
}

func NewClient(userID string, conn *websocket.Conn) *Client {
	return &Client{
		UserID: userID,
		Conn:   conn,
	}
}
