package handler

import (
	"institute/helpers"
	helper "institute/helpers"
	"strconv"

	"institute/features/realtime_chat"
	"institute/features/realtime_chat/dtos"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type controller struct {
	service realtime_chat.Usecase
}

func New(service realtime_chat.Usecase) realtime_chat.Handler {
	return &controller {
		service: service,
	}
}

var validate *validator.Validate

func (ctl *controller) GetRealtime_chats() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		pagination := dtos.Pagination{}
		ctx.Bind(&pagination)
		
		page := pagination.Page
		size := pagination.Size

		if page <= 0 || size <= 0 {
			return ctx.JSON(400, helper.Response("Please provide query `page` and `size` in number!"))
		}

		realtime_chats := ctl.service.FindAll(page, size)

		if realtime_chats == nil {
			return ctx.JSON(404, helper.Response("There is No Realtime_chats!"))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": realtime_chats,
		}))
	}
}


func (ctl *controller) Realtime_chatDetails() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		realtime_chatID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response(err.Error()))
		}

		realtime_chat := ctl.service.FindByID(realtime_chatID)

		if realtime_chat == nil {
			return ctx.JSON(404, helper.Response("Realtime_chat Not Found!"))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": realtime_chat,
		}))
	}
}

func (ctl *controller) CreateRealtime_chat() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		input := dtos.InputRealtime_chat{}

		ctx.Bind(&input)

		validate = validator.New(validator.WithRequiredStructEnabled())

		err := validate.Struct(input)

		if err != nil {
			errMap := helpers.ErrorMapValidation(err)
			return ctx.JSON(400, helper.Response("Bad Request!", map[string]any {
				"error": errMap,
			}))
		}

		realtime_chat := ctl.service.Create(input)

		if realtime_chat == nil {
			return ctx.JSON(500, helper.Response("Something went Wrong!", nil))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": realtime_chat,
		}))
	}
}

func (ctl *controller) UpdateRealtime_chat() echo.HandlerFunc {
	return func (ctx echo.Context) error {
		input := dtos.InputRealtime_chat{}

		realtime_chatID, errParam := strconv.Atoi(ctx.Param("id"))

		if errParam != nil {
			return ctx.JSON(400, helper.Response(errParam.Error()))
		}

		realtime_chat := ctl.service.FindByID(realtime_chatID)

		if realtime_chat == nil {
			return ctx.JSON(404, helper.Response("Realtime_chat Not Found!"))
		}
		
		ctx.Bind(&input)

		validate = validator.New(validator.WithRequiredStructEnabled())
		err := validate.Struct(input)

		if err != nil {
			errMap := helpers.ErrorMapValidation(err)
			return ctx.JSON(400, helper.Response("Bad Request!", map[string]any {
				"error": errMap,
			}))
		}

		update := ctl.service.Modify(input, realtime_chatID)

		if !update {
			return ctx.JSON(500, helper.Response("Something Went Wrong!"))
		}

		return ctx.JSON(200, helper.Response("Realtime_chat Success Updated!"))
	}
}

func (ctl *controller) DeleteRealtime_chat() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		realtime_chatID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response(err.Error()))
		}

		realtime_chat := ctl.service.FindByID(realtime_chatID)

		if realtime_chat == nil {
			return ctx.JSON(404, helper.Response("Realtime_chat Not Found!"))
		}

		delete := ctl.service.Remove(realtime_chatID)

		if !delete {
			return ctx.JSON(500, helper.Response("Something Went Wrong!"))
		}

		return ctx.JSON(200, helper.Response("Realtime_chat Success Deleted!", nil))
	}
}
