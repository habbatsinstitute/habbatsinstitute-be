package routes

import (
	"institute/features/product"

	"github.com/labstack/echo/v4"
)

func Products(e *echo.Echo, handler product.Handler) {
	products := e.Group("/products")

	products.GET("", handler.GetProducts())
	products.POST("", handler.CreateProduct())
	
	products.GET("/:id", handler.ProductDetails())
	products.PUT("/:id", handler.UpdateProduct())
	products.DELETE("/:id", handler.DeleteProduct())
}