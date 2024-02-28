package websocket

import (
	"encoding/base64"
	"fmt"

	"institute/utils/websocket/packet"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type Client struct {
	server  *Server
	handler *websocket.Conn
	rooms   map[int]*Room
	message chan *packet.Message
	sign    string
}

func NewClient(context echo.Context, server *Server, user int, role string) (string, *Client) {
	var (
		format = fmt.Sprintf("%s@%d", role, user)
		sign   = base64.RawStdEncoding.EncodeToString([]byte(format))
	)
	return format, &Client{
		server:  server,
		handler: NewProtocol().Switch(context),
		rooms:   make(map[int]*Room),
		message: make(chan *packet.Message),
		sign:    sign,
	}
}