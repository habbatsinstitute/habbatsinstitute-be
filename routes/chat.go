package routes

import (
	realtimechat "institute/features/realtime_chat"

	"github.com/labstack/echo/v4"
)

func Chats(e *echo.Echo, handler realtimechat.Handler) {
	e.GET("/chats/:id/:role_id", handler.Establish())

	rooms := e.Group("/rooms")
	rooms.GET("", handler.GetRooms())
}