package exception

import "errors"

type AuthException struct{}

func UserInvalid() error {
	return errors.New("user invalid")
}

func TokenInvalid() error {
	return errors.New("token invalid")
}

func CannotCreateSession() error {
	return errors.New("cannot create session")
}

func AuthenticationFailed() error {
	return errors.New("authentication failed")
}
