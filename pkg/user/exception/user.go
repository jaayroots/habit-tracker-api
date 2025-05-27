package exception

import "errors"

type UserException struct{}

func IsExistUser() error {
	return errors.New("User is already")
}

func CannotCreateUser() error {
	return errors.New("Cannot create user")
}

func CannotUpdateUser() error {
	return errors.New("Cannot update user")
}

func CannotDeleteUser() error {
	return errors.New("Cannot delete user")
}

func CannotFindUser() error {
	return errors.New("Cannot find user")
}

func NotFoundUser() error {
	return errors.New("Not found user")
}
