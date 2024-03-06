package routes

import (
	"institute/config"
	realtimechat "institute/features/realtime_chat"
	"institute/helpers"

	m "institute/middlewares"

	"github.com/labstack/echo/v4"
)

func Chats(e *echo.Echo, handler realtimechat.Handler, jwt helpers.JWTInterface, config config.ProgramConfig) {
	e.GET("/chats/:role/:id", handler.Establish())

	rooms := e.Group("/rooms")
	rooms.GET("", handler.GetRooms())
	rooms.GET("/details", handler.GetRoomBySenderId(), m.AuthorizeJWT(jwt, 3, config.SECRET))
	rooms.POST("/messages", handler.SaveChat(), m.AuthorizeJWT(jwt, 3, config.SECRET))
}