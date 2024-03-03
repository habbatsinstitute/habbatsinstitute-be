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
		
		role := ctx.Param("role_id")
		if role == "" {
			return ctx.JSON(http.StatusBadRequest, helpers.Response(err.Error()))
		}

		room, err := strconv.Atoi(ctx.Param("room"))
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, helpers.Response(err.Error()))
		}

		h.service.SocketEstablish(ctx, user, role, room)
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