package websocket

import (
	"institute/utils/websocket/packet"
)

type Room struct {
	clients map[*Client]bool
	join    chan *Client
	leave   chan *Client
	message chan *packet.Message
	sign    int
}