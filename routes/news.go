package routes

import (
	"institute/features/news"

	"github.com/labstack/echo/v4"
)

func Newss(e *echo.Echo, handler news.Handler) {
	newss := e.Group("/newss")

	newss.GET("", handler.GetNewss())
	newss.POST("", handler.CreateNews())
	
	newss.GET("/:id", handler.NewsDetails())
	newss.PUT("/:id", handler.UpdateNews())
	newss.DELETE("/:id", handler.DeleteNews())
	newss.GET("/category", handler.GetCategory())
}