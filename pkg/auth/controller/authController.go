package controller

import "github.com/labstack/echo/v4"

type AuthContoller interface {
	Register(pctx echo.Context) error
	Login(pctx echo.Context) error
	Logout(pctx echo.Context) error
	Refresh(pctx echo.Context) error
}
