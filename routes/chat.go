package routes

import (
	realtimechat "institute/features/realtimeChat"

	"github.com/labstack/echo/v4"
)

func Chats(e *echo.Echo, handler realtimechat.Handler) {
	e.GET("/chats/:user_id/:role_id/:room_id", handler.Establish())
}