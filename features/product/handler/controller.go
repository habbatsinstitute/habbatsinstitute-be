package handler

import (
	"institute/helpers"
	helper "institute/helpers"
	"strconv"

	"institute/features/product"
	"institute/features/product/dtos"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type controller struct {
	service product.Usecase
}

func New(service product.Usecase) product.Handler {
	return &controller {
		service: service,
	}
}

var validate *validator.Validate

func (ctl *controller) GetProducts() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		pagination := dtos.Pagination{}
		ctx.Bind(&pagination)
		
		page := pagination.Page
		size := pagination.Size

		if page <= 0 || size <= 0 {
			return ctx.JSON(400, helper.Response("Please provide query `page` and `size` in number!"))
		}

		products := ctl.service.FindAll(page, size)

		if products == nil {
			return ctx.JSON(404, helper.Response("There is No Products!"))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": products,
		}))
	}
}


func (ctl *controller) ProductDetails() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		productID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response(err.Error()))
		}

		product := ctl.service.FindByID(productID)

		if product == nil {
			return ctx.JSON(404, helper.Response("Product Not Found!"))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": product,
		}))
	}
}

func (ctl *controller) CreateProduct() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		input := dtos.InputProduct{}

		ctx.Bind(&input)

		validate = validator.New(validator.WithRequiredStructEnabled())

		err := validate.Struct(input)

		if err != nil {
			errMap := helpers.ErrorMapValidation(err)
			return ctx.JSON(400, helper.Response("Bad Request!", map[string]any {
				"error": errMap,
			}))
		}

		product := ctl.service.Create(input)

		if product == nil {
			return ctx.JSON(500, helper.Response("Something went Wrong!", nil))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": product,
		}))
	}
}

func (ctl *controller) UpdateProduct() echo.HandlerFunc {
	return func (ctx echo.Context) error {
		input := dtos.InputProduct{}

		productID, errParam := strconv.Atoi(ctx.Param("id"))

		if errParam != nil {
			return ctx.JSON(400, helper.Response(errParam.Error()))
		}

		product := ctl.service.FindByID(productID)

		if product == nil {
			return ctx.JSON(404, helper.Response("Product Not Found!"))
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

		update := ctl.service.Modify(input, productID)

		if !update {
			return ctx.JSON(500, helper.Response("Something Went Wrong!"))
		}

		return ctx.JSON(200, helper.Response("Product Success Updated!"))
	}
}

func (ctl *controller) DeleteProduct() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		productID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response(err.Error()))
		}

		product := ctl.service.FindByID(productID)

		if product == nil {
			return ctx.JSON(404, helper.Response("Product Not Found!"))
		}

		delete := ctl.service.Remove(productID)

		if !delete {
			return ctx.JSON(500, helper.Response("Something Went Wrong!"))
		}

		return ctx.JSON(200, helper.Response("Product Success Deleted!", nil))
	}
}
