package controller

import "github.com/labstack/echo/v4"

type CheckinContoller interface {
	Create(pctx echo.Context) error
	FindAll(pctx echo.Context) error
	Delete(pctx echo.Context) error
}
