package controller

import (
	"context"
	"net/http"
	"strings"

	"github.com/jaayroots/habit-tracker-api/custom"
	"github.com/jaayroots/habit-tracker-api/pctxkeys"
	"github.com/jaayroots/habit-tracker-api/utils"

	"github.com/labstack/echo/v4"

	_service "github.com/jaayroots/habit-tracker-api/service"
)

type middlewareContollerImpl struct {
	authService _service.AuthService
}

func NewMiddlewareControllerImpl(
	authService _service.AuthService,
) MiddlewareContoller {
	return &middlewareContollerImpl{
		authService: authService,
	}
}

type MiddlewareContoller interface {
	Authorizing(pctx echo.Context, next echo.HandlerFunc) error
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

	ctx := context.WithValue(pctx.Request().Context(), pctxkeys.ContextKeyUserID, loginRes.User.ID)
	ctx = context.WithValue(ctx, pctxkeys.ContextKeyLang, lang)

	req := pctx.Request().WithContext(ctx)

	pctx.SetRequest(req)

	return next(pctx)
}
