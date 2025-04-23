package usecase

import (
	"fmt"
	realtimechat "institute/features/realtimeChat"
	"institute/utils/websocket"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type service struct {
	socket *websocket.Server
}

func New(socket *websocket.Server) realtimechat.Usecase {
	return &service{
		socket: socket,
	}
}

func(svc *service) SocketEstablish(ctx echo.Context, userID int, roleID int, roomID int) {
	fmt.Printf("userid: %d", userID)
	fmt.Printf("roleid: %d", roleID)
	fmt.Printf("roomid: %d", roomID)
	sign := svc.socket.CreateClient(ctx, userID, roleID)
	ref := fmt.Sprintf("%d@%d", roleID, userID)

	if svc.socket.FindRoom(roomID) == nil {
		if roomID <= 0 {
			logrus.Errorf("Room %d gagal dibuat!", roomID)
		} else {
			logrus.Infof("Room %d berhasil dibuat!", roomID)
		}
		svc.socket.CreateRoom(roomID, ref)
	}else {
		svc.socket.JoinRoom(roomID, ref)
	}
	ctx.Set("ws.connect", sign)
}