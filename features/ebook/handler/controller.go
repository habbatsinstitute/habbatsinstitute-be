package handler

import (
	"errors"
	"fmt"
	"institute/helpers"
	helper "institute/helpers"
	"strconv"

	"institute/features/ebook"
	"institute/features/ebook/dtos"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type controller struct {
	service ebook.Usecase
}

func New(service ebook.Usecase) ebook.Handler {
	return &controller {
		service: service,
	}
}

var validate *validator.Validate

func (ctl *controller) GetEbooks() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		pagination := dtos.Pagination{}
		ctx.Bind(&pagination)
		
		page := pagination.Page
		size := pagination.Size

		if page <= 0 || size <= 0 {
			page = 1
			size = 5
		}

		ebooks := ctl.service.FindAll(page, size)

		if ebooks == nil {
			return ctx.JSON(404, helper.Response("There is No Ebooks!"))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": ebooks,
		}))
	}
}


func (ctl *controller) EbookDetails() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		ebookID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response(err.Error()))
		}

		ebook := ctl.service.FindByID(ebookID)

		if ebook == nil {
			return ctx.JSON(404, helper.Response("Ebook Not Found!"))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": ebook,
		}))
	}
}

func (ctl *controller) CreateEbook() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		input := dtos.InputEbook{}
		fileHeader, err := ctx.FormFile("ebook")
		thumbnailFileHeader, err := ctx.FormFile("thumbnail_book")

		ctx.Bind(&input)
		userID := ctx.Get("user_id")

		validate = validator.New(validator.WithRequiredStructEnabled())

		err = validate.Struct(input)

		if err != nil {
			errMap := helpers.ErrorMapValidation(err)
			return ctx.JSON(400, helpers.Response("Bad Request!", map[string]any{
				"error": errMap,
			}))
		}
		ebook, errMap, err := ctl.service.Create(input, userID.(int), fileHeader, thumbnailFileHeader)
		if errMap != nil {
			return ctx.JSON(400, helpers.Response("missing some data", map[string]any{
				"error": errMap,
			}))
		}
		if err != nil {
			return errors.New("failed to create")
		}
		if ebook == nil {
			return ctx.JSON(500, helpers.Response("something went wrong!", nil))
		}
		fmt.Println("data: ", ebook)
		return ctx.JSON(200, helpers.Response("succes", map[string]any{
			"data":ebook,
		}))
	}
}

func (ctl *controller) UpdateEbook() echo.HandlerFunc {
	return func (ctx echo.Context) error {
		input := dtos.InputEbook{}

		fileHeader, err := ctx.FormFile("ebook")

		if err  != nil {
			return ctx.JSON(400, helper.Response("Bad Request!", map[string]any {
				"error ebook": err,
			}))
		}
		thumbnailFileHeader, err := ctx.FormFile("thumbnail_book")

		// if err  != nil {
		// 	return ctx.JSON(400, helper.Response("Bad Request!", map[string]any {
		// 		"error thumbnail book": err,
		// 	}))
		// }

		ebookID, errParam := strconv.Atoi(ctx.Param("id"))

		if errParam != nil {
			return ctx.JSON(400, helper.Response(errParam.Error()))
		}

		ebook := ctl.service.FindByID(ebookID)

		if ebook == nil {
			return ctx.JSON(404, helper.Response("Ebook Not Found!"))
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

		update := ctl.service.Modify(input, ebookID, fileHeader, thumbnailFileHeader)

		if !update {
			return ctx.JSON(500, helper.Response("Something Went Wrong!"))
		}

		return ctx.JSON(200, helper.Response("Ebook Success Updated!"))
	}
}

func (ctl *controller) DeleteEbook() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		ebookID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response(err.Error()))
		}

		ebook := ctl.service.FindByID(ebookID)

		if ebook == nil {
			return ctx.JSON(404, helper.Response("Ebook Not Found!"))
		}

		delete := ctl.service.Remove(ebookID)

		if !delete {
			return ctx.JSON(500, helper.Response("Something Went Wrong!"))
		}

		return ctx.JSON(200, helper.Response("Ebook Success Deleted!", nil))
	}
}
