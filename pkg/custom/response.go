package custom

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ResponseFormat struct {
	Success       bool        `json:"success"`
	Code          int         `json:"code"`
	Message       string      `json:"message,omitempty"`
	Data          interface{} `json:"data,omitempty"`
	Errors        string      `json:"errors,omitempty"`
	IsNeedRefresh *bool       `json:"is_need_refresh,omitempty"`
}

func Response(pctx echo.Context, httpStatus int, data interface{}, message string, err error) error {
	status := true
	if httpStatus != http.StatusOK {
		status = false
	}

	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}

	var isNeedRefresh *bool
	if val, ok := pctx.Get("isTokenExpSoon").(bool); ok && val {
		trueVal := true
		isNeedRefresh = &trueVal
	}

	res := ResponseFormat{
		Success:       status,
		Code:          httpStatus,
		Message:       message,
		Data:          data,
		Errors:        errMsg,
		IsNeedRefresh: isNeedRefresh,
	}

	return pctx.JSON(httpStatus, res)
}
