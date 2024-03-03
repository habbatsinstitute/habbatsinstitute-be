package realtimechat

import (
	"github.com/labstack/echo/v4"
)

type Repository interface {
}

type Usecase interface {
	SocketEstablish(ctx echo.Context, userId int, role string, roomId int)
}

type Handler interface {
	Establish() echo.HandlerFunc
}