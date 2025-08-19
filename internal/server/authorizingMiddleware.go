package server

import (
	_controller "github.com/habit-tracker-api/controller"
	"github.com/labstack/echo/v4"
)

type authorizingMiddleware struct {
	middlewareContoller _controller.MiddlewareContoller
}

func (m *authorizingMiddleware) Authorizing(next echo.HandlerFunc) echo.HandlerFunc {
	return func(pctx echo.Context) error {
		return m.middlewareContoller.Authorizing(pctx, next)
	}
}
