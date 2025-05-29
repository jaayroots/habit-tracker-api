package service

import (
	_checkinModel "github.com/jaayroots/habit-tracker-api/pkg/checkin/model"
	"github.com/labstack/echo/v4"
)

type CheckinService interface {
	Create(pctx echo.Context, checkinReq *_checkinModel.CheckinReq) (*_checkinModel.CheckinRes, error)
	FindAll(pctx echo.Context, checkinSearchReq *_checkinModel.CheckinSearchReq) (*_checkinModel.CheckinSearchRes, error)
	Delete(pctx echo.Context, checkinID uint) (*_checkinModel.CheckinRes, error)
}
