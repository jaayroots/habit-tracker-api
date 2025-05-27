package controller

import (
	"context"
	"net/http"
	"strings"

	"github.com/jaayroots/go_base/pctxkeys"
	"github.com/jaayroots/go_base/pkg/custom"
	"github.com/jaayroots/go_base/utils"

	"github.com/labstack/echo/v4"

	_authService "github.com/jaayroots/go_base/pkg/auth/service"
)

type middlewareContollerImpl struct {
	authService _authService.AuthService
}

func NewMiddlewareControllerImpl(
	authService _authService.AuthService,
) MiddlewareContoller {
	return &middlewareContollerImpl{
		authService: authService,
	}
}

func (c *middlewareContollerImpl) Authorizing(pctx echo.Context, next echo.HandlerFunc) error {

	authHeader := pctx.Request().Header.Get("Authorization")
	if authHeader == "" {
		return custom.Response(pctx, http.StatusUnauthorized, nil, "", nil)
	}

	token := ""
	if strings.HasPrefix(authHeader, "Bearer ") {
		token = strings.TrimPrefix(authHeader, "Bearer ")
	}

	loginRes, isTokenExpSoon, err := c.authService.Authorizing(token)
	if err != nil {
		return custom.Response(pctx, http.StatusUnauthorized, nil, "", err)
	}

	pctx.Set("isTokenExpSoon", isTokenExpSoon)
	pctx.Set("user", loginRes.User)
	lang := utils.ValidateLangOrDefault(pctx)

	ctx := context.WithValue(pctx.Request().Context(), pctxkeys.ContextKeyUserID, uint(loginRes.User.ID))
	ctx = context.WithValue(ctx, pctxkeys.ContextKeyLang, lang)

	req := pctx.Request().WithContext(ctx)

	pctx.SetRequest(req)

	return next(pctx)
}
