package controller

import "github.com/labstack/echo/v4"

type MiddlewareContoller interface {
	Authorizing(pctx echo.Context, next echo.HandlerFunc) error
}
