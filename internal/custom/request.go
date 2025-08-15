package custom

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"sync"
)

type (
	EchoRequest interface {
		Build(obj any) error
	}

	customEchoRequest struct {
		ctx       echo.Context
		validator *validator.Validate
	}
)

var (
	once              sync.Once
	validatorInstanct *validator.Validate
)

func NewCustomEchoRequest(echoRequest echo.Context) EchoRequest {
	once.Do(func() {
		validatorInstanct = validator.New()
	})
	return &customEchoRequest{
		ctx:       echoRequest,
		validator: validatorInstanct,
	}
}

func (r *customEchoRequest) Build(obj any) error {
	if err := r.ctx.Bind(obj); err != nil {
		return err
	}
	if err := r.validator.Struct(obj); err != nil {
		return err
	}

	return nil
}
