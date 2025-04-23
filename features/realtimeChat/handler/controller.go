package handler

import (
	realtimechat "institute/features/realtimeChat"
	"institute/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type controller struct {
	service realtimechat.Usecase
}

func New(service realtimechat.Usecase) realtimechat.Handler {
	return &controller{
		service: service,
	}
}

func (c *controller) Establish() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// userID := ctx.QueryParam("user_id")
		// convUserID, _ := strconv.Atoi(userID)
		// roleID := ctx.QueryParam("role_id")
		// convRoleID, _ := strconv.Atoi(roleID)
		// roomID := ctx.QueryParam("room_id")
		// convRoomID, _ := strconv.Atoi(roomID)
		// fmt.Printf("userid: %s", userID)
		// fmt.Printf("roleid: %s", roleID)
		// fmt.Printf("roomid: %s", roomID)
		// fmt.Printf("after conv")
		// fmt.Printf("userid: %d", convUserID)
		// fmt.Printf("roleid: %d", convRoleID)
		// fmt.Printf("roomid: %d", convRoomID)
		convUserID := 123
		convRoleID := 456
		convRoomID := 111
		c.service.SocketEstablish(ctx, convUserID, convRoleID, convRoomID)
		
		if msg := ctx.Get("ws.client.error"); msg != nil {
			logrus.Infof("[ws.establish]: %d@%d not found", convUserID,convRoleID)
			response := helpers.ResponseError{
				Status:  false,
				Message: msg.(string),
			}
			return ctx.JSON(http.StatusNotFound, response)
		}
		if client := ctx.Get("ws.connect"); client != nil {
			logrus.Infof("[ws.establish]: client@%s connected", client)
		}
		return nil
	}
}