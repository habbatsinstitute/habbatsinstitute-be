package realtimechat

import (
	"institute/features/realtime_chat/dtos"
	"time"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	TimeStamp() time.Time
	SaveChat(questionNReply Chat, userID int, recipientID int) error
	SelectByID(userID int) *User
}

type Usecase interface {
	SocketEstablish(ctx echo.Context, userId int, role string, roomId int)
	GetRooms() []dtos.RoomRes
	GetRoomBySenderId(senderId int) *dtos.RoomRes
	SaveChat(ctx echo.Context, req dtos.Request, userID int) *dtos.ChatRes
}

type Handler interface {
	Establish() echo.HandlerFunc
	GetRooms() echo.HandlerFunc
	GetRoomBySenderId() echo.HandlerFunc
	SaveChat() echo.HandlerFunc
}