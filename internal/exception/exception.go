package exception

import (
	"errors"
	_exceptionType "github.com/jaayroots/habit-tracker-api/enums/exception"
)

type Exception struct{}

func Handle(message string, exceptionType _exceptionType.ExceptionType) error {
	switch exceptionType {
	case _exceptionType.Info, _exceptionType.Warning:
		return errors.New(message)
	case _exceptionType.Error, _exceptionType.Fatal:
		return errors.New(message)
	default:
		return errors.New(message)
	}
}
