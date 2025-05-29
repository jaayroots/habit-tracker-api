package repository

import (
	_checkinModel "github.com/jaayroots/habit-tracker-api/pkg/checkin/model"

	"github.com/jaayroots/habit-tracker-api/entities"
	"github.com/labstack/echo/v4"
)

type CheckinRepository interface {
	Create(pctx echo.Context, checkin *entities.Checkin) (*entities.Checkin, error)
	FindAll(pctx echo.Context, checkinSearchReq *_checkinModel.CheckinSearchReq) ([]*entities.Checkin, int, error)
	FindByID(pctx echo.Context, checkinID uint) (*entities.Checkin, error)
	Delete(pctx echo.Context, checkinID uint) (*entities.Checkin, error)
	GroupByHabitIDcheckin(pctx echo.Context, habitIDs []uint) ([]*_checkinModel.GroupByHabitIDcheckin, error)
}
