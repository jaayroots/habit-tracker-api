package controller

import (
	"net/http"

	_checkinService "github.com/jaayroots/habit-tracker-api/pkg/checkin/service"
	"github.com/jaayroots/habit-tracker-api/pkg/custom"
	"github.com/jaayroots/habit-tracker-api/utils"
	"github.com/labstack/echo/v4"

	_checkinModel "github.com/jaayroots/habit-tracker-api/pkg/checkin/model"
)

type checkinContollerImpl struct {
	checkinService _checkinService.CheckinService
}

func NewCheckinControllerImpl(
	checkinService _checkinService.CheckinService,
) CheckinContoller {
	return &checkinContollerImpl{
		checkinService,
	}
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
