package routes

import (
	"institute/config"
	"institute/features/auth"
	"institute/helpers"

	m "institute/middlewares"

	"github.com/labstack/echo/v4"
)

func Auth(e *echo.Echo, handler auth.Handler, jwt helpers.JWTInterface, config config.ProgramConfig){
	auth := e.Group("/auth")
	auth.POST("/register", handler.RegisterUser(), m.AuthorizeJWT(jwt, 2, config.SECRET)  )
	auth.POST("/login", handler.Login())
	auth.POST("/refresh-jwt", handler.RefreshJWT(), m.AuthorizeJWT(jwt, 3, config.SECRET) )
}