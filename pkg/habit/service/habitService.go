package service

import (
	_habitModel "github.com/jaayroots/habit-tracker-api/pkg/habit/model"
	"github.com/labstack/echo/v4"
)

type HabitService interface {
	Create(pctx echo.Context, habitReq *_habitModel.HabitReq) (*_habitModel.HabitRes, error)
	FindByID(pctx echo.Context, habitID uint) (*_habitModel.HabitRes, error)
	Update(pctx echo.Context, habitID uint, habitReq *_habitModel.HabitReq) (*_habitModel.HabitRes, error)
	Delete(pctx echo.Context, habitID uint) (*_habitModel.HabitRes, error)
	FindAll(pctx echo.Context, habitSearchReq *_habitModel.HabitSearchReq) (*_habitModel.HabitSearchRes, error)
}
