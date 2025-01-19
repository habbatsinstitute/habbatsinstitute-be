package gemini

import (
	"institute/features/gemini/dtos"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	InsertInformation(ctx echo.Context, payload dtos.InformationBodyRequest) (id int, err error)
	EditInformation(ctx echo.Context, payload dtos.InformationBodyRequest, id int64) (isSuccess bool, err error)
	DeleteInformation(ctx echo.Context, id int64) (success bool, err error)
	GetInformations(ctx echo.Context) (informations []Information, err error)
	GetInformationByID(ctx echo.Context, id int64) (information Information, err error)
}

type Usecase interface {
	InsertInformation(ctx echo.Context, payload dtos.InformationBodyRequest) (information dtos.InformationResponse, err error)
	EditInformation(ctx echo.Context, payload dtos.InformationBodyRequest, id int64) (information dtos.InformationResponse, err error)
	DeleteInformation(ctx echo.Context, id int64) (information dtos.InformationResponse, err error)
	GetInformations(ctx echo.Context) (informations []dtos.InformationResponse, err error)
	GetInformationByID(ctx echo.Context, id int64) (information dtos.InformationResponse, err error)

	GetChatResponse(eCtx echo.Context, question string) (answer string, err error)
}

type Handler interface {
	InsertInformation(ctx echo.Context) error
	EditInformation(ctx echo.Context) error
	DeleteInformation(ctx echo.Context) error
	GetInformationByID(ctx echo.Context) error
	GetInformation(ctx echo.Context) error

	GetChatResponse(eCtx echo.Context) error
}
