package routes

import (
	"institute/config"
	"institute/features/ebook"
	"institute/helpers"

	m "institute/middlewares"

	"github.com/labstack/echo/v4"
)

func Ebooks(e *echo.Echo, handler ebook.Handler, jwt helpers.JWTInterface, config config.ProgramConfig) {
	ebooks := e.Group("/ebooks")

	ebooks.GET("", handler.GetEbooks(), m.AuthorizeJWT(jwt, 3, config.SECRET))
	ebooks.POST("", handler.CreateEbook(), m.AuthorizeJWT(jwt, 3, config.SECRET))
	
	ebooks.GET("/:id", handler.EbookDetails())
	ebooks.PUT("/:id", handler.UpdateEbook(), m.AuthorizeJWT(jwt, 3, config.SECRET))
	ebooks.DELETE("/:id", handler.DeleteEbook(),)
}