package exception

import "errors"

type AuthException struct{}

func UserInvalid() error {
	return errors.New("User invalid")
}

func TokenInvalid() error {
	return errors.New("Token invalid")
}

func CannotCreateSession() error {
	return errors.New("Cannot create session")
}

func AuthenticationFailed() error {
	return errors.New("Authentication failed")
}
