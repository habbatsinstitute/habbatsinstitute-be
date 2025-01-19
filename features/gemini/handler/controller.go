package handler

import (
	"institute/helpers"
	"net/http"
	"strconv"

	"institute/features/gemini"
	"institute/features/gemini/dtos"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type controller struct {
	service gemini.Usecase
}

func New(service gemini.Usecase) gemini.Handler {
	return &controller {
		service: service,
	}
}

var validate *validator.Validate

func (c *controller) InsertInformation(ctx echo.Context) error {
	var body dtos.InformationBodyRequest
	if err := ctx.Bind(&body); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"success": false,
			"message": helpers.ErrBadPayloadRequest.Error(),
		})
	}

	information, err := c.service.InsertInformation(ctx, body)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, echo.Map{
		"success": true,
		"data":    information,
	})
}

func (c *controller) EditInformation(ctx echo.Context) error {
	idString := ctx.Param("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"success": false,
			"message": helpers.ErrBadPayloadRequest.Error(),
		})
	}

	var body dtos.InformationBodyRequest
	if err := ctx.Bind(&body); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"success": false,
			"message": helpers.ErrBadPayloadRequest.Error(),
		})
	}

	information, err := c.service.EditInformation(ctx, body, id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"success": true,
		"data":    information,
	})
}

func (c *controller) DeleteInformation(ctx echo.Context) error {
	idString := ctx.Param("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"success": false,
			"message": helpers.ErrBadPayloadRequest.Error(),
		})
	}

	information, err := c.service.DeleteInformation(ctx, id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"success": true,
		"data":    information,
	})
}

func (c *controller) GetInformationByID(ctx echo.Context) error {
	idString := ctx.Param("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"success": false,
			"message": helpers.ErrBadPayloadRequest.Error(),
		})
	}

	information, err := c.service.GetInformationByID(ctx, id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"success": true,
		"data":    information,
	})
}

func (c *controller) GetInformation(ctx echo.Context) error {
	informations, err := c.service.GetInformations(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"success": true,
		"data":    informations,
	})
}

func (c *controller) GetChatResponse(eCtx echo.Context) error {
    var body dtos.ChatRequestPayload
    if err := eCtx.Bind(&body); err != nil {
        return eCtx.JSON(http.StatusBadRequest, echo.Map{
            "success": false,
            "message": helpers.ErrBadPayloadRequest.Error(),
        })
    }

    answer, err := c.service.GetChatResponse(eCtx, body.Question)
    if err != nil {
        return eCtx.JSON(http.StatusInternalServerError, echo.Map{
            "success": false,
            "message": err.Error(),
        })
    }

    return eCtx.JSON(http.StatusOK, echo.Map{
        "success": true,
        "data": dtos.ChatResponse{
            Question: body.Question,
            Answer:   answer,
        },
    })
}
