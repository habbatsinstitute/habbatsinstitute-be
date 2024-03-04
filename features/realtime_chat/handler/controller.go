package handler

import (
	realtimechat "institute/features/realtime_chat"
	"institute/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Chathandler struct {
	service realtimechat.Usecase
}

func New(service realtimechat.Usecase) realtimechat.Handler{
	return &Chathandler{
		service: service,
	}
}

func (h *Chathandler) Establish() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, helpers.Response(err.Error()))
		}
		
		role, _ := strconv.Atoi(ctx.Param("role_id"))
		if role == 0 {
			return ctx.JSON(http.StatusBadRequest, helpers.Response(err.Error()))
		}

		roomId, _ := strconv.Atoi(ctx.QueryParam("room_id"))
		if roomId == 0 {
			return ctx.JSON(http.StatusBadRequest, helpers.Response(err.Error()))
		}

		h.service.SocketEstablish(ctx, user, role, roomId)
		if msg := ctx.Get("ws.client.error"); msg != nil {
			logrus.Infof("[ws.establish]: %v not found", user)
			response := helpers.Response("not found", map[string]any{
				"data": msg,
			})
			return ctx.JSON(http.StatusNotFound, response)
		}
		if client := ctx.Get("ws.connect"); client != nil {
			logrus.Infof("[ws.establish]: client@%s connected", client)
		}
		return nil
	}
}

func (h *Chathandler) GetRooms() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		
		rooms := h.service.GetRooms()
		if rooms == nil {
			return ctx.JSON(404, helpers.Response("There is No Rooms!"))
		}

		return ctx.JSON(200, helpers.Response("Success!", map[string]any {
			"data": rooms,
		}))
	}
}
