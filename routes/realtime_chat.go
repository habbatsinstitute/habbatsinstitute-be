package routes

import (
	"institute/features/realtime_chat"

	"github.com/labstack/echo/v4"
)

func Realtime_chats(e *echo.Echo, handler realtime_chat.Handler) {
	realtime_chats := e.Group("/realtime_chats")

	realtime_chats.GET("", handler.GetRealtime_chats())
	realtime_chats.POST("", handler.CreateRealtime_chat())
	
	realtime_chats.GET("/:id", handler.Realtime_chatDetails())
	realtime_chats.PUT("/:id", handler.UpdateRealtime_chat())
	realtime_chats.DELETE("/:id", handler.DeleteRealtime_chat())
}