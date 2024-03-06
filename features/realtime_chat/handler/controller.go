package handler

import (
	realtimechat "institute/features/realtime_chat"
	"institute/features/realtime_chat/dtos"
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
		
		role := ctx.Param("role")
		if role == "" {
			return ctx.JSON(http.StatusBadRequest, helpers.Response(err.Error()))
		}

		roomId, _ := strconv.Atoi(ctx.QueryParam("room_id"))
		// if roomId == 0 {
		// 	return ctx.JSON(http.StatusBadRequest, helpers.Response(err.Error()))
		// }

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

func (h *Chathandler) GetRoomBySenderId() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		senderId := ctx.Get("user_id").(int)

		if senderId == 0 {
			return ctx.JSON(400, helpers.Response("User ID not found!"))
		}

		room := h.service.GetRoomBySenderId(senderId)

		if room == nil {
			return ctx.JSON(404, helpers.Response("Room Not Found!"))
		}

		return ctx.JSON(200, helpers.Response("Success!", map[string]any {
			"data": room,
		}))
	}
}

func (h *Chathandler) SaveChat() echo.HandlerFunc{
	return func(ctx echo.Context) error {
		input := dtos.Request{}
		
		ctx.Bind(&input)
		userID := 0

		if ctx.Get("user_id") != nil {
			userID = ctx.Get("user_id").(int)
		}

		chat := h.service.SaveChat(ctx, input, userID)
		
		if chat == nil {
			return ctx.JSON(500, helpers.Response("Something went Wrong!", nil))
		}

		return ctx.JSON(200, helpers.Response("succes!", map[string]any{
			"data": chat,
		}))
	}
}
