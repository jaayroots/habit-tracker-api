package exception

import "errors"

type HabitException struct{}

func IsExistHabit() error {
	return errors.New("habit is already")
}

func CannotCreateHabit() error {
	return errors.New("cannot create habit")
}

func CannotUpdateHabit() error {
	return errors.New("cannot update habit")
}

func CannotDeleteHabit() error {
	return errors.New("cannot delete habit")
}

func CannotFindHabit() error {
	return errors.New("cannot find habit")
}

func NotFoundHabit() error {
	return errors.New("not found habit")
}

func FrequencyInvalid() error {
	return errors.New("frequency invalid")
}
