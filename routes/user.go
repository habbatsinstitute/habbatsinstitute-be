package routes

import (
	"institute/config"
	"institute/features/user"
	"institute/helpers"

	m "institute/middlewares"

	"github.com/labstack/echo/v4"
)

func Users(e *echo.Echo, handler user.Handler, jwt helpers.JWTInterface, config config.ProgramConfig) {
	users := e.Group("/users")

	users.GET("", handler.GetUsers(), m.AuthorizeJWT(jwt, 2, config.SECRET))
	users.GET("/me", handler.MyProfile(), m.AuthorizeJWT(jwt, 3, config.SECRET))
	
	users.GET("/:id", handler.UserDetails(), m.AuthorizeJWT(jwt, 3, config.SECRET))
	users.DELETE("/:id", handler.DeleteUser(), m.AuthorizeJWT(jwt, 2, config.SECRET))

	users.PUT("/update/:id", handler.UpdateExpiryAccount(), m.AuthorizeJWT(jwt, 3, config.SECRET))
}