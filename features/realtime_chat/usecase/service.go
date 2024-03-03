package usecase

import (
	"fmt"
	realtimechat "institute/features/realtime_chat"
	"institute/features/user"
	"institute/utils/websocket"

	"github.com/labstack/echo/v4"
)

type ChatService struct {
	socket *websocket.Server
	user user.Repository
}

func New(socket *websocket.Server, user user.Repository) realtimechat.Usecase{
	return &ChatService{
		socket: socket,
		user: user,
	}
}

func (s *ChatService) SocketEstablish(ctx echo.Context, userId int, role string, roomId int) {
	user := s.user.SelectByID(userId)
	if user == nil {
		ctx.Set("ws.client.error", "user not found")
		return
	}
	sign := s.socket.CreateClient(ctx, userId, role, roomId)
	ref := fmt.Sprintf("%s@%d@%d", role, userId, roomId) 
	if s.socket.FindRoom(roomId) == nil {
		s.socket.CreateRoom(roomId, ref)
	} else {
		s.socket.JoinRoom(roomId, ref)
	}
	ctx.Set("ws.connect", sign)
}