package exception

import "errors"

type UserException struct{}

func IsExistEmail() error {
	return errors.New("email is already")
}

func CannotCreateUser() error {
	return errors.New("cannot create user")
}

func CannotUpdateUser() error {
	return errors.New("cannot update user")
}

func CannotDeleteUser() error {
	return errors.New("cannot delete user")
}

func CannotFindUser() error {
	return errors.New("cannot find user")
}

func NotFoundUser() error {
	return errors.New("not found user")
}
