package exception

import "errors"

type CheckinException struct{}

func CannotCreateCheckin() error {
	return errors.New("Cannot create checkin")
}

func CannotDeleteCheckin() error { 
	return errors.New("Cannot delete checkin")
}

func NotFoundCheckin() error { 
	return errors.New("Not found checkin")
}


