package exception

import "errors"

type CheckinException struct{}

func CannotCreateCheckin() error {
	return errors.New("Cannot create checkin")
}
