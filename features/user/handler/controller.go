package handler

import (
	helpers "institute/helpers"
	"strconv"

	"institute/features/user"
	"institute/features/user/dtos"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type controller struct {
	service user.Usecase
}

func New(service user.Usecase) user.Handler {
	return &controller {
		service: service,
	}
}

var validate *validator.Validate

func (ctl *controller) GetUsers() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		pagination := dtos.Pagination{}
		ctx.Bind(&pagination)
		
		if pagination.Page < 1 || pagination.Size < 1 {
			pagination.Page = 1
			pagination.Size = 10
		}

		page := pagination.Page
		size := pagination.Size

		users := ctl.service.FindAll(page, size)

		if users == nil {
			return ctx.JSON(404, helpers.Response("There is No Users!"))
		}

		return ctx.JSON(200, helpers.Response("Success!", map[string]any {
			"data": users,
		}))
	}
}


func (ctl *controller) UserDetails() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		userID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helpers.Response(err.Error()))
		}

		user := ctl.service.FindByID(userID)

		if user == nil {
			return ctx.JSON(404, helpers.Response("User Not Found!"))
		}

		return ctx.JSON(200, helpers.Response("Success!", map[string]any {
			"data": user,
		}))
	}
}

func (ctl *controller) UpdateUser() echo.HandlerFunc {
	return func (ctx echo.Context) error {
		input := dtos.InputUser{}

		userID, errParam := strconv.Atoi(ctx.Param("id"))

		if errParam != nil {
			return ctx.JSON(400, helpers.Response(errParam.Error()))
		}

		user := ctl.service.FindByID(userID)

		if user == nil {
			return ctx.JSON(404, helpers.Response("User Not Found!"))
		}
		
		ctx.Bind(&input)

		validate = validator.New(validator.WithRequiredStructEnabled())
		err := validate.Struct(input)

		if err != nil {
			errMap := helpers.ErrorMapValidation(err)
			return ctx.JSON(400, helpers.Response("Bad Request!", map[string]any {
				"error": errMap,
			}))
		}

		update := ctl.service.Modify(input, userID)

		if !update {
			return ctx.JSON(500, helpers.Response("Something Went Wrong!"))
		}

		return ctx.JSON(200, helpers.Response("User Success Updated!"))
	}
}

func (ctl *controller) UpdateExpiryAccount() echo.HandlerFunc {
	return func (ctx echo.Context) error {
		input := dtos.UpdateUser{}

		userID, errParam := strconv.Atoi(ctx.Param("id"))

		if errParam != nil {
			return ctx.JSON(400, helpers.Response(errParam.Error()))
		}

		user := ctl.service.FindByID(userID)

		if user == nil {
			return ctx.JSON(404, helpers.Response("User Not Found!"))
		}
		
		ctx.Bind(&input)

		validate = validator.New(validator.WithRequiredStructEnabled())
		err := validate.Struct(input)

		if err != nil {
			errMap := helpers.ErrorMapValidation(err)
			return ctx.JSON(400, helpers.Response("Bad Request!", map[string]any {
				"error": errMap,
			}))
		}

		update := ctl.service.ModifyUser(input, userID)

		if !update {
			return ctx.JSON(500, helpers.Response("Something Went Wrong!"))
		}

		return ctx.JSON(200, helpers.Response("User Success Updated!"))
	}
}

func (ctl *controller) DeleteUser() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		userID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helpers.Response(err.Error()))
		}

		user := ctl.service.FindByID(userID)

		if user == nil {
			return ctx.JSON(404, helpers.Response("User Not Found!"))
		}

		delete := ctl.service.Remove(userID)

		if !delete {
			return ctx.JSON(500, helpers.Response("Something Went Wrong!"))
		}

		return ctx.JSON(200, helpers.Response("User Success Deleted!", nil))
	}
}

func (ctl *controller) MyProfile() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userID := ctx.Get("user_id").(int)

		user := ctl.service.MyProfile(userID)
		if user == nil {
			return ctx.JSON(404, helpers.Response("user not found"))
		}
		
		return ctx.JSON(200, helpers.Response("succes", map[string]any{
			"data": user,
		}))
	}
}