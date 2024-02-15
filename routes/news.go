package routes

import (
	"institute/config"
	"institute/features/news"
	"institute/helpers"
	m "institute/middlewares"

	"github.com/labstack/echo/v4"
)

func Newss(e *echo.Echo, handler news.Handler, jwt helpers.JWTInterface, config config.ProgramConfig) {
	newss := e.Group("/news")

	newss.POST("", handler.CreateNews(), m.AuthorizeJWT(jwt, 2, config.SECRET))
	
	newss.GET("/category", handler.GetCategory())
	newss.GET("/:id", handler.NewsDetails())
	newss.GET("/searching", handler.SearchNewsByTitle())
	newss.GET("", handler.GetNews())

	newss.PATCH("/:id", handler.UpdateNews(), m.AuthorizeJWT(jwt, 2, config.SECRET))
	newss.DELETE("/:id", handler.DeleteNews(), m.AuthorizeJWT(jwt, 2, config.SECRET))
}