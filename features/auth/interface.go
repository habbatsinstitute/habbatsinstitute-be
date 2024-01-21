package auth

import (
	"institute/features/auth/dtos"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Register(newUser *User) (*User, error)
	Login(email string) (*User, error)
	SelectByUsername(username string) (*User, error)
}

type Usecase interface {
	Register(newData dtos.InputUser) (*dtos.ResUser, []string, error)
	Login(data dtos.RequestLogin) (*dtos.LoginResponse, []string, error)
	RefreshJWT(jwt dtos.RefreshJWT) (*dtos.ResJWT, error)
}

type Handler interface {
	RegisterUser() echo.HandlerFunc
	Login() echo.HandlerFunc
	RefreshJWT() echo.HandlerFunc
}