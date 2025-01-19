package routes

import (
	"institute/features/gemini"

	"github.com/labstack/echo/v4"
)

func Geminis(e *echo.Echo, handler gemini.Handler) {
	geminis := e.Group("/gemini")

	geminis.GET("", handler.GetInformation)
	geminis.POST("", handler.InsertInformation)
	
	geminis.GET("/:id", handler.GetInformationByID)
	geminis.PUT("/:id", handler.EditInformation)
	geminis.DELETE("/:id", handler.DeleteInformation)

	geminis.POST("/chat", handler.GetChatResponse)
}