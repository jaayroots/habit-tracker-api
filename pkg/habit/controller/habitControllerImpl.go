package controller

import (
	"net/http"

	"github.com/jaayroots/habit-tracker-api/pkg/custom"
	_habitService "github.com/jaayroots/habit-tracker-api/pkg/habit/service"
	"github.com/labstack/echo/v4"

	_habitModel "github.com/jaayroots/habit-tracker-api/pkg/habit/model"
	_utils "github.com/jaayroots/habit-tracker-api/utils"
)

type habitContollerImpl struct {
	habitService _habitService.HabitService
}

func NewHabitControllerImpl(
	habitService _habitService.HabitService,
) HabitContoller {
	return &habitContollerImpl{
		habitService,
	}
}

func (c *habitContollerImpl) Create(pctx echo.Context) error {

	habitReq := new(_habitModel.HabitReq)
	customerEchoRequest := custom.NewCustomEchoRequest(pctx)
	if err := customerEchoRequest.Build(habitReq); err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, "Invalid request", err)
	}

	_, err := c.habitService.Create(pctx, habitReq)
	if err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, err.Error(), nil)
	}

	return custom.Response(pctx, http.StatusOK, nil, "", nil)

}

func (c *habitContollerImpl) FindByID(pctx echo.Context) error {

	habitID, err := _utils.StrToUint(pctx.Param("habitID"))
	if err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, "Invalid habitID", nil)
	}

	habit, err := c.habitService.FindByID(pctx, habitID)
	if err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, err.Error(), nil)
	}

	return custom.Response(pctx, http.StatusOK, habit, "", nil)

}

func (c *habitContollerImpl) Update(pctx echo.Context) error {

	habitID, err := _utils.StrToUint(pctx.Param("habitID"))
	if err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, "Invalid habitID", nil)
	}

	habitReq := new(_habitModel.HabitReq)
	customerEchoRequest := custom.NewCustomEchoRequest(pctx)
	if err := customerEchoRequest.Build(habitReq); err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, "Invalid request", err)
	}

	_, err = c.habitService.Update(pctx, habitID, habitReq)
	if err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, err.Error(), nil)
	}

	return custom.Response(pctx, http.StatusOK, nil, "", nil)

}

func (c *habitContollerImpl) Delete(pctx echo.Context) error {

	habitID, err := _utils.StrToUint(pctx.Param("habitID"))
	if err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, "Invalid habitID", nil)
	}

	_, err = c.habitService.Delete(pctx, habitID)
	if err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, err.Error(), nil)
	}

	return custom.Response(pctx, http.StatusOK, nil, "", nil)

}

func (c *habitContollerImpl) FindAll(pctx echo.Context) error {

	habitSearchReq := new(_habitModel.HabitSearchReq)
	customerEchoRequest := custom.NewCustomEchoRequest(pctx)
	if err := customerEchoRequest.Build(habitSearchReq); err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, "Invalid request", err)
	}

	habitSearch, err := c.habitService.FindAll(pctx, habitSearchReq)
	if err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, err.Error(), nil)
	}

	return custom.Response(pctx, http.StatusOK, habitSearch, "", nil)

}
