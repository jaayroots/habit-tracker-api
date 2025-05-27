package controller

import (
	"net/http"

	"github.com/jaayroots/go_base/pkg/custom"

	"github.com/labstack/echo/v4"

	_authModel "github.com/jaayroots/go_base/pkg/auth/model"
	_authService "github.com/jaayroots/go_base/pkg/auth/service"
	_userModel "github.com/jaayroots/go_base/pkg/user/model"
)

type authContollerImpl struct {
	authService _authService.AuthService
}

func NewAuthControllerImpl(
	authService _authService.AuthService,
) AuthContoller {
	return &authContollerImpl{
		authService: authService,
	}
}

func (c *authContollerImpl) Register(pctx echo.Context) error {

	createReq := new(_userModel.UserReq)

	customerEchoRequest := custom.NewCustomEchoRequest(pctx)
	if err := customerEchoRequest.Build(createReq); err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, "Invalid request", err)
	}

	user, err := c.authService.Register(createReq)
	if err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, "", err)
	}

	return custom.Response(pctx, http.StatusOK, user, "Register successful", nil)

}

func (c *authContollerImpl) Login(pctx echo.Context) error {

	loginReq := new(_authModel.LoginReq)

	customerEchoRequest := custom.NewCustomEchoRequest(pctx)
	if err := customerEchoRequest.Build(loginReq); err != nil {
		return custom.Response(pctx, http.StatusBadRequest, nil, "Invalid request", err)
	}

	loginReq.IpAddress = pctx.RealIP()
	token, err := c.authService.Login(loginReq)
	if err != nil {
		return custom.Response(pctx, http.StatusUnauthorized, nil, "", err)
	}

	return custom.Response(pctx, http.StatusOK, token, "Login successful", nil)

}

func (c *authContollerImpl) Logout(pctx echo.Context) error {

	val := pctx.Get("user")
	user, ok := val.(*_userModel.UserRes)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "")
	}

	err := c.authService.Logout(uint(user.ID))
	if err != nil {
		return custom.Response(pctx, http.StatusUnauthorized, nil, "", err)
	}

	return custom.Response(pctx, http.StatusOK, "", "Logout successful", nil)

}

func (c *authContollerImpl) Refresh(pctx echo.Context) error {

	val := pctx.Get("user")
	user, ok := val.(*_userModel.UserRes)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "")
	}

	ipAddress := pctx.RealIP()
	token, err := c.authService.Refreash(ipAddress, uint(user.ID))
	if err != nil {
		return custom.Response(pctx, http.StatusUnauthorized, nil, "", err)
	}

	return custom.Response(pctx, http.StatusOK, token, "Refresh successful", nil)
}
