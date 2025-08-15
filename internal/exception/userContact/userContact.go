package exception

import "errors"

type UserContactException struct{}

func IsExistEmail() error {
	return errors.New("email is already")
}

func CannotCreateUserContact() error {
	return errors.New("cannot create user contact")
}

func CannotUpdateUserContact() error {
	return errors.New("cannot update user contact")
}

func CannotDeleteUserContact() error {
	return errors.New("cannot delete user contact")
}

func CannotFindUserContact() error {
	return errors.New("cannot find user contact")
}

func NotFoundUserContact() error {
	return errors.New("not found user contact")
}
