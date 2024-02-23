package realtime_chat

import (
	"institute/features/realtime_chat/dtos"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Paginate(page, size int) []Realtime_chat
	Insert(newRealtime_chat Realtime_chat) int64
	SelectByID(realtime_chatID int) *Realtime_chat
	Update(realtime_chat Realtime_chat) int64
	DeleteByID(realtime_chatID int) int64
}

type Usecase interface {
	FindAll(page, size int) []dtos.ResRealtime_chat
	FindByID(realtime_chatID int) *dtos.ResRealtime_chat
	Create(newRealtime_chat dtos.InputRealtime_chat) *dtos.ResRealtime_chat
	Modify(realtime_chatData dtos.InputRealtime_chat, realtime_chatID int) bool
	Remove(realtime_chatID int) bool
}

type Handler interface {
	GetRealtime_chats() echo.HandlerFunc
	Realtime_chatDetails() echo.HandlerFunc
	CreateRealtime_chat() echo.HandlerFunc
	UpdateRealtime_chat() echo.HandlerFunc
	DeleteRealtime_chat() echo.HandlerFunc
}
