package exception

import "errors"

type CheckinException struct{}

func CannotCreateCheckin() error {
	return errors.New("cannot create checkin")
}

func CannotDeleteCheckin() error {
	return errors.New("cannot delete checkin")
}

func NotFoundCheckin() error {
	return errors.New("not found checkin")
}
