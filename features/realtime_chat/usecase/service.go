package usecase

import (
	"fmt"
	realtimechat "institute/features/realtime_chat"
	"institute/features/realtime_chat/dtos"
	"institute/helpers"
	"institute/utils/websocket"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type ChatService struct {
	socket *websocket.Server
	model realtimechat.Repository
	generator helpers.GeneratorInterface
}

func New(socket *websocket.Server, model realtimechat.Repository, generator helpers.GeneratorInterface) realtimechat.Usecase{
	return &ChatService{
		socket: socket,
		model: model,
		generator: generator,
	}
}

func (s *ChatService) SocketEstablish(ctx echo.Context, userId int, role string, roomId int) {
	user := s.model.SelectByID(userId)
	if user == nil {
		ctx.Set("ws.client.error", "user not found")
		return
	}
	
	sign := s.socket.CreateClient(ctx, userId, role)
	ref := fmt.Sprintf("%s@%d", role, userId) 
	if s.socket.FindRoom(roomId) == nil {
		roomId := s.generator.GenerateRandomID()
		s.socket.CreateRoom(roomId, *user, ref)
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
			SenderName: room.SenderName,
			SenderId: room.SenderId,
		})	
	}

	return rooms
}

func (s *ChatService) GetRoomBySenderId(senderId int) *dtos.RoomRes {
	var room = s.socket.Rooms[senderId]

	fmt.Println("isi rooms", s.socket.Rooms)
	fmt.Println("isi room =", room)

	if room == nil {
		return nil
	}

	var res = &dtos.RoomRes{
		RoomId: room.ID,
		SenderName: room.SenderName,
		SenderId: room.SenderId,
	}

	return res
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

	resChat := dtos.ChatRes{}
	errRes := smapping.FillStruct(&resChat, smapping.MapFields(req))
	if errRes != nil {
		log.Error(errRes)
		return nil
	}
	return &resChat
}