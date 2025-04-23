package realtimechat

import "github.com/labstack/echo/v4"

type Repository interface {
}

type Usecase interface {
	SocketEstablish(ctx echo.Context, userID int, roleID int, roomID int)
}

type Handler interface {
	Establish() echo.HandlerFunc
}
