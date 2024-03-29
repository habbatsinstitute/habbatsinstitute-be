package handler

import (
	"errors"
	"institute/helpers"
	helper "institute/helpers"
	"strconv"

	"institute/features/news"
	"institute/features/news/dtos"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type controller struct {
	service news.Usecase
}

func New(service news.Usecase) news.Handler {
	return &controller {
		service: service,
	}
}

var validate *validator.Validate

func (ctl *controller) GetNews() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		pagination := dtos.Pagination{}
		ctx.Bind(&pagination)

		if pagination.Page < 1 || pagination.Size < 1 {
			pagination.Page = 1
			pagination.Size = 5
		}
		
		page := pagination.Page
		size := pagination.Size

		newss, totalData := ctl.service.FindAll(page, size)

		if newss == nil {
			return ctx.JSON(404, helper.Response("There is No Newss!", map[string]any{
				"data": []string{},
			}))
		}

		paginationResponse := helpers.PaginationResponse(page, size, int(totalData))

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": newss,
			"pagination":paginationResponse,
		}))
	}
}

func (ctl *controller) GetTopnews() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		pagination := dtos.Pagination{}

		pagination.Size = 3
		size := pagination.Size

		newss := ctl.service.GetTopNews(size)

		if newss == nil {
			return ctx.JSON(404, helper.Response("there is no news", map[string]any{
				"data": []string{},
			}))
		}
		return ctx.JSON(200, helper.Response("succes", map[string]any{
			"data": newss,
		}))
	}
}


func (ctl *controller) NewsDetails() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		newsID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response(err.Error()))
		}

		news := ctl.service.FindByID(newsID)

		if news.ID < 0 {
			return ctx.JSON(404, helper.Response("News Not Found!", map[string]any{
				"data":news,
			}))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": news,
		}))
	}
}

func (ctl *controller) CreateNews() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		input := dtos.InputNews{}
		filHeader, err := ctx.FormFile("images")

		ctx.Bind(&input)

		userID := ctx.Get("user_id")

		validate = validator.New(validator.WithRequiredStructEnabled())

		err = validate.Struct(input)

		if err != nil {
			errMap := helpers.ErrorMapValidation(err)
			return ctx.JSON(400, helper.Response("Bad Request!", map[string]any {
				"error": errMap,
			}))
		}

		news,errMap, err := ctl.service.Create(input,userID.(int),filHeader)
		
		if errMap != nil {
			return ctx.JSON(400, helpers.Response("missing some data", map[string]any{
				"error": errMap,
			}))
		}

		if err != nil {
			return errors.New("failed to create")
		}
		if news == nil {
			return ctx.JSON(500, helper.Response("something went wrong!", nil))
		}
		return ctx.JSON(200, helper.Response("succes", map[string]any {
			"data": news,
		}))
	}
}

func (ctl *controller) UpdateNews() echo.HandlerFunc {
	return func (ctx echo.Context) error {
		input := dtos.InputNews{}

		fileHeader, err := ctx.FormFile("images")

		newsID, errParam := strconv.Atoi(ctx.Param("id"))

		if errParam != nil {
			return ctx.JSON(400, helper.Response(errParam.Error()))
		}

		news := ctl.service.FindByID(newsID)

		if news == nil {
			return ctx.JSON(404, helper.Response("News Not Found!"))
		}
		
		ctx.Bind(&input)

		validate = validator.New(validator.WithRequiredStructEnabled())
		err = validate.Struct(input)

		if err != nil {
			errMap := helpers.ErrorMapValidation(err)
			return ctx.JSON(400, helper.Response("Bad Request!", map[string]any {
				"error": errMap,
			}))
		}

		update := ctl.service.Modify(input, newsID, fileHeader)

		if !update {
			return ctx.JSON(500, helper.Response("Something Went Wrong!"))
		}

		return ctx.JSON(200, helper.Response("News Success Updated!"))
	}
}

func (ctl *controller) DeleteNews() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		newsID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response(err.Error()))
		}

		news := ctl.service.FindByID(newsID)

		if news == nil {
			return ctx.JSON(404, helper.Response("News Not Found!"))
		}

		delete := ctl.service.Remove(newsID)

		if !delete {
			return ctx.JSON(500, helper.Response("Something Went Wrong!"))
		}

		return ctx.JSON(200, helper.Response("News Success Deleted!", map[string]any{
			"data":news,
		}))
	}
}


func (ctl *controller) GetCategory() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		Category, err := ctl.service.FindAllCategory()

		if err != nil {
			return ctx.JSON(500, helpers.Response(err.Error()))
		}

		return ctx.JSON(200, helpers.Response("succes", map[string]any {
			"data":Category,
		}))
	}
}

func (ctl *controller) SearchNewsByTitle() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		title := ctx.QueryParam("title")
		titles, err := ctl.service.SearchNews(title)
		if err != nil {
			return ctx.JSON(400, helpers.Response("not found"))
		}
		return ctx.JSON(200, helpers.Response("succes", map[string]any{
			"data": titles,
		}))
	}
}