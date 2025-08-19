package controller

import (
	"net/http"

	"github.com/habit-tracker-api/custom"
	_service "github.com/habit-tracker-api/service"
	"github.com/habit-tracker-api/utils"
	"github.com/labstack/echo/v4"

	_checkinModel "github.com/habit-tracker-api/model/checkin"
)

type checkinContollerImpl struct {
	checkinService _service.CheckinService
}

func NewCheckinControllerImpl(
	checkinService _service.CheckinService,
) CheckinContoller {
	return &checkinContollerImpl{
		checkinService,
	}
}

type CheckinContoller interface {
	Create(pctx echo.Context) error
	FindAll(pctx echo.Context) error
	Delete(pctx echo.Context) error
}

func (c *checkinContollerImpl) Create(pctx echo.Context) error {

	checkinReq := new(_checkinModel.CheckinReq)
	customerEchoRequest := custom.NewCustomEchoRequest(pctx)
	if err := customerEchoRequest.Build(checkinReq); err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, "Invalid request", err)
	}

	_, err := c.checkinService.Create(pctx, checkinReq)
	if err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, err.Error(), nil)
	}

	return custom.Response(pctx, http.StatusOK, nil, "", nil)

}

func (c *checkinContollerImpl) FindAll(pctx echo.Context) error {

	checkinSearchReq := new(_checkinModel.CheckinSearchReq)
	customerEchoRequest := custom.NewCustomEchoRequest(pctx)
	if err := customerEchoRequest.Build(checkinSearchReq); err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, "Invalid request", err)
	}

	checkinSearch, err := c.checkinService.FindAll(pctx, checkinSearchReq)
	if err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, err.Error(), nil)
	}

	return custom.Response(pctx, http.StatusOK, checkinSearch, "", nil)

}

func (c *checkinContollerImpl) Delete(pctx echo.Context) error {

	checkinID, err := utils.StrToUint(pctx.Param("checkinID"))
	if err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, "Invalid checkinID", nil)
	}

	_, err = c.checkinService.Delete(pctx, checkinID)
	if err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, err.Error(), nil)
	}

	return custom.Response(pctx, http.StatusOK, nil, "", nil)

}
