package handler

import (
	"institute/helpers"
	helper "institute/helpers"
	"strconv"

	"institute/features/item"
	"institute/features/item/dtos"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type controller struct {
	service item.Usecase
}

func New(service item.Usecase) item.Handler {
	return &controller {
		service: service,
	}
}

var validate *validator.Validate

func (ctl *controller) GetItems() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		pagination := dtos.Pagination{}
		ctx.Bind(&pagination)
		
		page := pagination.Page
		size := pagination.Size

		if page <= 0 || size <= 0 {
			return ctx.JSON(400, helper.Response("Please provide query `page` and `size` in number!"))
		}

		items := ctl.service.FindAll(page, size)

		if items == nil {
			return ctx.JSON(404, helper.Response("There is No Items!"))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": items,
		}))
	}
}


func (ctl *controller) ItemDetails() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		itemID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response(err.Error()))
		}

		item := ctl.service.FindByID(itemID)

		if item == nil {
			return ctx.JSON(404, helper.Response("Item Not Found!"))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": item,
		}))
	}
}

func (ctl *controller) CreateItem() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		input := dtos.InputItem{}

		ctx.Bind(&input)

		validate = validator.New(validator.WithRequiredStructEnabled())

		err := validate.Struct(input)

		if err != nil {
			errMap := helpers.ErrorMapValidation(err)
			return ctx.JSON(400, helper.Response("Bad Request!", map[string]any {
				"error": errMap,
			}))
		}

		item := ctl.service.Create(input)

		if item == nil {
			return ctx.JSON(500, helper.Response("Something went Wrong!", nil))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": item,
		}))
	}
}

func (ctl *controller) UpdateItem() echo.HandlerFunc {
	return func (ctx echo.Context) error {
		input := dtos.InputItem{}

		itemID, errParam := strconv.Atoi(ctx.Param("id"))

		if errParam != nil {
			return ctx.JSON(400, helper.Response(errParam.Error()))
		}

		item := ctl.service.FindByID(itemID)

		if item == nil {
			return ctx.JSON(404, helper.Response("Item Not Found!"))
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

		update := ctl.service.Modify(input, itemID)

		if !update {
			return ctx.JSON(500, helper.Response("Something Went Wrong!"))
		}

		return ctx.JSON(200, helper.Response("Item Success Updated!"))
	}
}

func (ctl *controller) DeleteItem() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		itemID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response(err.Error()))
		}

		item := ctl.service.FindByID(itemID)

		if item == nil {
			return ctx.JSON(404, helper.Response("Item Not Found!"))
		}

		delete := ctl.service.Remove(itemID)

		if !delete {
			return ctx.JSON(500, helper.Response("Something Went Wrong!"))
		}

		return ctx.JSON(200, helper.Response("Item Success Deleted!", nil))
	}
}
