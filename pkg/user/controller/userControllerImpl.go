package controller

import (
	"net/http"

	"github.com/jaayroots/habit-tracker-api/pkg/custom"
	"github.com/jaayroots/habit-tracker-api/utils"
	"github.com/labstack/echo/v4"

	_userModel "github.com/jaayroots/habit-tracker-api/pkg/user/model"
	_userService "github.com/jaayroots/habit-tracker-api/pkg/user/service"
)

type userContollerImpl struct {
	userService _userService.UserService
}

func NewUserControllerImpl(
	userService _userService.UserService,
) UserContoller {
	return &userContollerImpl{
		userService,
	}
}

func (c *userContollerImpl) FindByID(pctx echo.Context) error {

	userID, err := utils.StrToUint(pctx.Param("userID"))
	if err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, "Invalid userID", nil)
	}

	user, err := c.userService.FindByID(uint(userID))
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

	userID, err := utils.StrToUint(pctx.Param("userID"))
	if err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, "Invalid userID", nil)
	}

	err = c.userService.Update(uint(userID), updateReq)
	if err != nil {
		return custom.Response(pctx, http.StatusInternalServerError, nil, "", err)
	}

	return custom.Response(pctx, http.StatusOK, "", "", nil)

}

func (c *userContollerImpl) Delete(pctx echo.Context) error {

	userID, err := utils.StrToUint(pctx.Param("userID"))
	if err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, "Invalid userID", nil)
	}

	err = c.userService.Delete(uint(userID))
	if err != nil {
		return custom.Response(pctx, http.StatusInternalServerError, nil, "", err)
	}

	return custom.Response(pctx, http.StatusOK, "", "", nil)

}
