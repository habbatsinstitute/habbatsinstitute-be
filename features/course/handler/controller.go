package handler

import (
	"institute/helpers"
	helper "institute/helpers"
	"strconv"

	"institute/features/course"
	"institute/features/course/dtos"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type controller struct {
	service course.Usecase
}

func New(service course.Usecase) course.Handler {
	return &controller {
		service: service,
	}
}

var validate *validator.Validate

func (ctl *controller) GetCourses() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		pagination := dtos.Pagination{}
		ctx.Bind(&pagination)

		if pagination.Page < 1 || pagination.PageSize < 1 {
			pagination.Page = 1
			pagination.PageSize = 20
		}

		search := dtos.Search{}
		ctx.Bind(&search)
		
		page := pagination.Page
		pageSize := pagination.PageSize

		courses, totalData := ctl.service.FindAll(page, pageSize, search)

		if courses == nil {
			return ctx.JSON(404, helper.Response("There is No Courses!"))
		}

		paginationResponse := helpers.PaginationResponse(page, pageSize, int(totalData))

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": courses,
			"pagination":paginationResponse,
		}))
	}
}


func (ctl *controller) CourseDetails() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		courseID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response(err.Error()))
		}

		course := ctl.service.FindByID(courseID)

		if course == nil {
			return ctx.JSON(404, helper.Response("Course Not Found!"))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": course,
		}))
	}
}

func (ctl *controller) CreateCourse() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		input := dtos.InputCourse{}
		fileHeader, err := ctx.FormFile("media_file")

		ctx.Bind(&input)

		validate = validator.New(validator.WithRequiredStructEnabled())

		err = validate.Struct(input)

		if err != nil {
			errMap := helpers.ErrorMapValidation(err)
			return ctx.JSON(400, helper.Response("Bad Request!", map[string]any {
				"error": errMap,
			}))
		}

		course, err := ctl.service.Create(input, fileHeader)

		if course == nil {
			return ctx.JSON(500, helper.Response("Something went Wrong!", nil))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": course,
		}))
	}
}

func (ctl *controller) UpdateCourse() echo.HandlerFunc {
	return func (ctx echo.Context) error {
		input := dtos.InputCourse{}

		courseID, errParam := strconv.Atoi(ctx.Param("id"))

		if errParam != nil {
			return ctx.JSON(400, helper.Response(errParam.Error()))
		}

		course := ctl.service.FindByID(courseID)

		if course == nil {
			return ctx.JSON(404, helper.Response("Course Not Found!"))
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

		update := ctl.service.Modify(input, courseID)

		if !update {
			return ctx.JSON(500, helper.Response("Something Went Wrong!"))
		}

		return ctx.JSON(200, helper.Response("Course Success Updated!"))
	}
}

func (ctl *controller) DeleteCourse() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		courseID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response(err.Error()))
		}

		course := ctl.service.FindByID(courseID)

		if course == nil {
			return ctx.JSON(404, helper.Response("Course Not Found!"))
		}

		delete := ctl.service.Remove(courseID)

		if !delete {
			return ctx.JSON(500, helper.Response("Something Went Wrong!"))
		}

		return ctx.JSON(200, helper.Response("Course Success Deleted!", nil))
	}
}
