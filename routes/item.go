package routes

import (
	"institute/features/item"

	"github.com/labstack/echo/v4"
)

func Items(e *echo.Echo, handler item.Handler) {
	items := e.Group("/items")

	items.GET("", handler.GetItems())
	items.POST("", handler.CreateItem())
	
	items.GET("/:id", handler.ItemDetails())
	items.PUT("/:id", handler.UpdateItem())
	items.DELETE("/:id", handler.DeleteItem())
}