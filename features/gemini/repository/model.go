package repository

import (
	"errors"
	"institute/features/gemini"
	"institute/features/gemini/dtos"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	ErrDB       = errors.New("database error")
	ErrNotFound = errors.New("record not found")
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) InsertInformation(ctx echo.Context, payload dtos.InformationBodyRequest) (id int, err error) {
	information := gemini.Information{
		Question: payload.Question,
		Answer:   payload.Answer,
	}

	if err = r.db.WithContext(ctx.Request().Context()).Create(&information).Error; err != nil {
		return 0, ErrDB
	}

	return information.ID, nil
}

func (r *Repository) EditInformation(ctx echo.Context, payload dtos.InformationBodyRequest, id int64) (isSuccess bool, err error) {
	information := gemini.Information{}

	if err = r.db.WithContext(ctx.Request().Context()).First(&information, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, ErrNotFound
		}
		return false, ErrDB
	}

	information.Question = payload.Question
	information.Answer = payload.Answer

	if err = r.db.WithContext(ctx.Request().Context()).Save(&information).Error; err != nil {
		return false, ErrDB
	}

	return true, nil
}

func (r *Repository) DeleteInformation(ctx echo.Context, id int64) (success bool, err error) {
	if err = r.db.WithContext(ctx.Request().Context()).Delete(&gemini.Information{}, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, ErrNotFound
		}
		return false, ErrDB
	}

	return true, nil
}

func (r *Repository) GetInformations(ctx echo.Context) (informations []gemini.Information, err error) {
	if err = r.db.WithContext(ctx.Request().Context()).Find(&informations).Error; err != nil {
		return nil, ErrDB
	}
	return informations, nil
}

func (r *Repository) GetInformationByID(ctx echo.Context, id int64) (information gemini.Information, err error) {
	if err = r.db.WithContext(ctx.Request().Context()).First(&information, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gemini.Information{}, ErrNotFound
		}
		return gemini.Information{}, ErrDB
	}
	return information, nil
}
