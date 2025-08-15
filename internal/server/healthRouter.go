package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *echoServer) healthRouter() {
	router := s.app.Group("/health")
	router.GET("", s.healthCheck)
}

func (s *echoServer) healthCheck(pctx echo.Context) error {
	return pctx.String(http.StatusOK, "OK")
}
