package routes

import (
	"institute/config"
	"institute/features/product"
	"institute/helpers"

	m "institute/middlewares"

	"github.com/labstack/echo/v4"
)

func Products(e *echo.Echo, handler product.Handler, jwt helpers.JWTInterface, config config.ProgramConfig) {
	products := e.Group("/products")

	products.GET("", handler.GetProducts(), m.AuthorizeJWT(jwt, 3, config.SECRET))
	products.POST("", handler.CreateProduct(), m.AuthorizeJWT(jwt, 2, config.SECRET))
	
	products.GET("/:id", handler.ProductDetails())
	products.PUT("/:id", handler.UpdateProduct())
	products.DELETE("/:id", handler.DeleteProduct())
}