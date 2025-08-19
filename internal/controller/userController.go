package controller

import (
	"net/http"

	"github.com/habit-tracker-api/custom"
	"github.com/habit-tracker-api/utils"
	"github.com/labstack/echo/v4"

	_userModel "github.com/habit-tracker-api/model/user"
	_service "github.com/habit-tracker-api/service"
)

type userContollerImpl struct {
	userService _service.UserService
}

func NewUserController(
	userService _service.UserService,
) UserContoller {
	return &userContollerImpl{
		userService,
	}
}

type UserContoller interface {
	FindByID(pctx echo.Context) error
	Update(pctx echo.Context) error
	Delete(pctx echo.Context) error
}

func (c *userContollerImpl) FindByID(pctx echo.Context) error {

	userID, err := utils.StrToUUID(pctx.Param("userID"))
	if err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, "Invalid userID", nil)
	}

	user, err := c.userService.FindByID(userID)
	if err != nil {
		return custom.Response(pctx, http.StatusNotFound, nil, "", err)
	}

	return custom.Response(pctx, http.StatusOK, user, "", nil)

}

func (c *userContollerImpl) Update(pctx echo.Context) error {

	updateReq := new(_userModel.UserUpdateReq)

	customerEchoRequest := custom.NewCustomEchoRequest(pctx)
	if err := customerEchoRequest.Build(updateReq); err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, "Invalid request", err)
	}

	userID, err := utils.StrToUUID(pctx.Param("userID"))
	if err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, "Invalid userID", nil)
	}

	err = c.userService.Update(userID, updateReq)
	if err != nil {
		return custom.Response(pctx, http.StatusInternalServerError, nil, "", err)
	}

	return custom.Response(pctx, http.StatusOK, "", "", nil)

}

func (c *userContollerImpl) Delete(pctx echo.Context) error {

	userID, err := utils.StrToUUID(pctx.Param("userID"))
	if err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, "Invalid userID", nil)
	}

	err = c.userService.Delete(userID)
	if err != nil {
		return custom.Response(pctx, http.StatusInternalServerError, nil, "", err)
	}

	return custom.Response(pctx, http.StatusOK, "", "", nil)

}
