package routes

import (
	"institute/config"
	"institute/features/chatbot"
	"institute/helpers"
	m "institute/middlewares"

	"github.com/labstack/echo/v4"
)

func Chatbots(e *echo.Echo, handler chatbot.Handler, jwt helpers.JWTInterface, config config.ProgramConfig) {
	chatbots := e.Group("/chatbots")

	chatbots.GET("", handler.GetChatHistory())
	chatbots.POST("", handler.SendQuestion())
	
	chatbots.DELETE("", handler.DeleteChatHistory(), m.AuthorizeJWT(jwt, 3, config.SECRET))
}
