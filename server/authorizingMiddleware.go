package server

import (
	_authController "github.com/jaayroots/go_base/pkg/auth/controller"
	"github.com/labstack/echo/v4"
)

type authorizingMiddleware struct {
	middlewareContoller _authController.MiddlewareContoller
}

func (m *authorizingMiddleware) Authorizing(next echo.HandlerFunc) echo.HandlerFunc {
	return func(pctx echo.Context) error {
		return m.middlewareContoller.Authorizing(pctx, next)
	}
}
