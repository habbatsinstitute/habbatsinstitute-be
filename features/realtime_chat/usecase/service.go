package usecase

import (
	"fmt"
	realtimechat "institute/features/realtime_chat"
	"institute/features/realtime_chat/dtos"
	"institute/utils/websocket"

	"github.com/labstack/echo/v4"
)

type ChatService struct {
	socket *websocket.Server
	model realtimechat.Repository
}

func New(socket *websocket.Server, model realtimechat.Repository) realtimechat.Usecase{
	return &ChatService{
		socket: socket,
		model: model,
	}
}

func (s *ChatService) SocketEstablish(ctx echo.Context, userId int, role int, roomId int) {
	user := s.model.SelectByID(userId)
	if user == nil {
		ctx.Set("ws.client.error", "user not found")
		return
	}
	sign := s.socket.CreateClient(ctx, userId, role)
	ref := fmt.Sprintf("%d@%d", role, userId) 
	if s.socket.FindRoom(roomId) == nil {
		s.socket.CreateRoom(roomId, roomId, ref)
	} else {
		s.socket.JoinRoom(roomId, ref)
	}
	ctx.Set("ws.connect", sign)
}

func (s *ChatService) GetRooms() []dtos.RoomRes{
	var rooms []dtos.RoomRes

	for _, room := range s.socket.Rooms{
		rooms = append(rooms, dtos.RoomRes{
			RoomId: room.ID,
		})	
	}

	return rooms
}

func (s *ChatService) SaveChat(ctx echo.Context, req dtos.Request, userID int) *dtos.ChatRes {
	timeStamp := s.model.TimeStamp()
	data := realtimechat.Chat{
		Text: req.Text,
		Blob: req.Blob,
		Timestamp: timeStamp,
	}

	if err := s.model.SaveChat(data, userID, req.RecipientID); err != nil {
		return nil
	}

	return &dtos.ChatRes{
		SenderID: userID,
		Text: data.Text,
		Blob: data.Blob,
		Timestamp: data.Timestamp,
	}
}