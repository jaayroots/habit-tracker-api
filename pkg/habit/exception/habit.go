package exception

import "errors"

type HabitException struct{}

func IsExistHabit() error {
	return errors.New("Habit is already")
}

func CannotCreateHabit() error {
	return errors.New("Cannot create habit")
}

func CannotUpdateHabit() error {
	return errors.New("Cannot update habit")
}

func CannotDeleteHabit() error {
	return errors.New("Cannot delete habit")
}

func CannotFindHabit() error {
	return errors.New("Cannot find habit")
}

func NotFoundHabit() error {
	return errors.New("Not found habit")
}

func FrequencyInvalid() error {
	return errors.New("Frequency invalid")
}